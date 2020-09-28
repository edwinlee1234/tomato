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
  </div>
</template>

<script>

export default {
  name: 'Home',
  data: function() {
		return {
      timeVal: 1500,
      timeInterval: null,
      isStart: false,
      customMinute: 5,
      startTimeVal: 1500,
		};
  },

  components: {
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
  }
}
</script>

<style scoped lang="scss">
* {
  text-align: center;
}

$pageWith: 600px;
%box {
    border: solid 1px;
    border-radius: 5px;
}

#time_block {
  width: $pageWith;
  margin: 0 auto;
  padding-bottom: 20px;
  @extend %box;

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
</style>