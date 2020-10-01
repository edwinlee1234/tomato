<template>
  <div class="home">
    <div id="time_block">
      <b-container id="select_timer">
        <b-row>
          <b-col id="long_break" v-on:click="longBreak">
            <b-button variant="outline-success">Long Break</b-button>
          </b-col>
          <b-col id="shot_break" v-on:click="shotBreak">
            <b-button variant="outline-success">Short Break</b-button>
          </b-col>
          <b-col id="custom">
            <b-button variant="outline-success">Custom</b-button>
          </b-col>
        </b-row>
      </b-container>
      <div id="time">
        {{ paddingzero(Math.floor(timeVal/60), 2) }}:{{ paddingzero(timeVal % 60, 2) }}
      </div>
      <b-button class="btn" v-on:click="start" v-if="!isStart" variant="outline-primary">Start</b-button>
      <b-button class="btn" v-on:click="stop" v-if="isStart" variant="outline-danger">Stop</b-button>
      <b-button class="btn" v-on:click="reset" v-if="!isStart">Reset</b-button>

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
  </div>
</template>

<script>
import { mapState } from 'vuex'
import TaskList from '@/components/TaskList'

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
		};
  },

  components: {
    "TaskList": TaskList
  },

  computed: {
    ...mapState({
      userInfo: state => state.user.info
    }),
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
      this.timeVal = this.timeVal - 1
      if (this.timeVal === 0) {
        this.finish()
      }
    },

    finish() {
      this.stop()
    },
    
    stop() {
      clearInterval(this.timeInterval)

      this.isStart = false
    },

    reset() {
      this.timeVal = this.startTimeVal
    },

    shotBreak() {
      this.stop()

      this.timeVal = 900
      this.startTimeVal = this.timeVal
    },

    longBreak() {
      this.stop()

      this.timeVal = 1500
      this.startTimeVal = this.timeVal
    },

    custom()  {
      if (this.customMinute > 60) {
        this.customMinute = 60
      }

      this.timeVal = this.customMinute * 60
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
  }
}
</script>

<style scoped lang="scss">
 @import '@/assets/css/variables.scss';

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
    font-size: 50px;
    padding-top: 30px;
    padding-bottom: 30px;
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
</style>