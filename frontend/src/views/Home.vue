<template>
  <div class="home">
    <div id="time_block">
      <b-container id="select_timer">
        <b-row>
          <b-col id="long_break" v-on:click="longBreak">
            <b-button variant="outline-light">Long Break</b-button>
          </b-col>
          <b-col id="shot_break" v-on:click="shotBreak">
            <b-button variant="outline-light">Short Break</b-button>
          </b-col>
          <b-col id="custom">
            <b-button variant="outline-light">Custom</b-button>
          </b-col>
        </b-row>
      </b-container>
      <div id="time">
        <radial-progress-bar :diameter="200"
                            :completed-steps="runTime"
                            :total-steps="startTimeVal">
          <p>{{ paddingzero(Math.floor(timeVal/60), 2) }}:{{ paddingzero(timeVal % 60, 2) }}</p>
        </radial-progress-bar>
        
      </div>
      <b-button class="btn" v-on:click="start" v-if="!isStart" variant="outline-light">Start</b-button>
      <b-button class="btn" v-on:click="stop" v-if="isStart" variant="danger">Stop</b-button>
      <b-button class="btn" v-on:click="reset" v-if="!isStart" variant="outline-light">Reset</b-button>

      <b-popover ref="popover" target="custom" title="Minute">
        <b-form-input v-model="customMinute" placeholder="Minute"></b-form-input>
        <br>
        <b-icon class="h5 mb-2" icon="check" v-on:click="custom">OK</b-icon>
        <b-icon class="h5 mb-2" icon="x" v-on:click="popupClose">X</b-icon>
      </b-popover>
    </div>

    <br>

    <div id="task_list" v-if="userInfo.isLogin">
      <div v-if="!taskItem">
        <h4>Choose a task</h4>
        <hr>
        <h4><b-icon class="plus_btn" v-on:click="showSelectTaskListModal()" icon="plus"></b-icon></h4>
      </div>
      <div v-else>
        <h4>Task</h4>
        <hr>
        <b-list-group>
          <b-list-group-item>{{ taskItem.groupName }} - {{ taskItem.name }} 
            <span class="task_btn">
              <b-icon v-on:click="showSelectTaskListModal()" icon="pencil-square"></b-icon>
              <b-icon v-on:click="clearSelectedTask()" icon="trash"></b-icon>
            </span>
          </b-list-group-item>
        </b-list-group>
      </div>
    </div>

    <TaskList ref="TaskList" @input="setTaskItem"></TaskList>

    <br>

    <div id="today_record" v-if="userInfo.isLogin">
      <div>
        <h4>Today</h4>
        <hr>
        <b-list-group>
          <b-list-group-item v-for="(record,index) in todayRecords" :key="index">
            <!--created_timestamp減掉spend_time才是開始任務的時候!-->
            {{ convertTimestampToDate(record.created_timestamp - record.spend_time) }}  {{ record.group_name }} - {{ record.name }}  {{ Math.floor(record.spend_time/60) }}  Minutes
          </b-list-group-item>
        </b-list-group>
      </div>
    </div>
  </div>

</template>

<script>
import { mapState } from 'vuex'
import TaskList from '@/components/TaskList'
import tool from '@/lib/tool'
import RadialProgressBar from 'vue-radial-progress'

export default {
  name: 'Home',
  data: function() {
		return {
      timeVal: 1500,
      timeInterval: null,
      isStart: false,
      customMinute: 5,
      startTimeVal: 1500,
      taskItem: null,
      todayRecords: [],
      runTime: 0,
		};
  },

  mixins: [tool],

  components: {
    "TaskList": TaskList,
    "RadialProgressBar": RadialProgressBar,
  },

  computed: {
    ...mapState({
      userInfo: state => state.user.info
    }),
  },

  watch: {
    userInfo: function () {
      if (this.userInfo.isLogin === true) {
        this.fetchTodayRecords()
      }
    },
  },

  mounted: function () {
    if (this.userInfo.isLogin === false) {
      return
    }
    this.fetchTodayRecords()
  },

  methods: {
    start() {
      if (this.timeVal === 0) {
        return
      }

      this.timeInterval = setInterval(this.countDown, 1000)
      this.isStart = true
    },

    countDown() {
      this.timeVal--
      this.runTime++

      if (this.timeVal === 0) {
        this.finish()
      }
    },
    
    stop() {
      clearInterval(this.timeInterval)

      this.isStart = false
    },

    reset() {
      this.timeVal = this.startTimeVal
      this.runTime = 0
    },

    shotBreak() {
      this.stop()

      this.timeVal = 900
      this.runTime = 0
      this.startTimeVal = this.timeVal
    },

    longBreak() {
      this.stop()

      this.timeVal = 1500
      this.runTime = 0
      this.startTimeVal = this.timeVal
    },

    custom()  {
      this.stop()
      if (this.customMinute > 60) {
        this.customMinute = 60
      }

      this.timeVal = this.customMinute * 60
      this.runTime = 0
      this.startTimeVal = this.timeVal
      this.popupClose()
    },

    paddingzero(num, length) {
      return (Array(length).join("0") + num).slice(-length)
    },

    popupClose() {
      this.$refs.popover.$emit('close')
    },

    showSelectTaskListModal() {
      this.$refs.TaskList.showSelectTaskListModal()
    },

    setTaskItem(payload) {
      this.taskItem = payload

      this.stop()
      this.reset()
    },

    clearSelectedTask() {
      this.taskItem = null

      this.stop()
      this.reset()
    },

    convertTimestampToDate(timestamp) {
      let date = new Date(timestamp * 1000);
      let hours = date.getHours()
      let minutes = "0" + date.getMinutes()
      let seconds = "0" + date.getSeconds()

      return  hours + ':' + minutes.substr(-2) + ':' + seconds.substr(-2)
    },

    finish() {
      this.stop()

      if (this.taskItem === null) {
        return
      }

      let self = this
      this.axios.post(this.APIURL + "/api/pomo", {
        "id": this.taskItem.id,
        "time": this.startTimeVal,
      })
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.fetchTodayRecords()
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    fetchTodayRecords() {
      let self = this
      this.axios.get(this.APIURL + "/api/day/record")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.todayRecords = resp.data.data.records
      })
      .catch(function (error) {
        self.Alert(error)
      })
    },
  }
}
</script>

<style scoped lang="scss">
 @import '@/assets/css/variables.scss';

.b-icon {
  cursor: pointer;
}

.list-group-item {
  background-color: $mainColor;
}

.home {
  hr {
    border-top: solid 1px $fontColor;
  }
}

#time_block {
  width: $pageWith;
  margin: 0 auto;
  padding-bottom: 20px;
  @extend %box;
  text-align: center;

  #select_timer {
    padding-top: 30px;
    padding-bottom: 30px;
    width: 500px;
  }

  #time {
    margin: 0 auto;
    margin-bottom: 20px;
    width: 200px;
    p {
      font-size: 40px;
      margin-top: 1rem;
      margin-bottom: 1rem;
    }
  }

  .btn {
    margin: 5px;
    cursor: pointer;  
  }
}

#task_list {
  width: $pageWith;
  margin: 0 auto;
  padding: 15px;
  text-align: left;
  @extend %box;

  .plus_btn {
    cursor: pointer;
  }

  .task_btn {
    cursor: pointer;
    position: relative;
    float:right;
    
    .b-icon {
      margin-left: 8px;
    }
  }
}

#today_record {
  width: $pageWith;
  margin: 0 auto;
  text-align: left;
  padding: 15px;
  @extend %box;
}
</style>