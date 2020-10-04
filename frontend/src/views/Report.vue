<template>
  <div id="report">
    <b-container> 
      <b-row>
        <b-col class="box">
          <ve-pie
            :data="pieData"
            :settings="pieChartSettings"
            :events="pieChartEvents">
          </ve-pie>
          <p v-on:click="initPieChart" class="back_btn">Back</p>
        </b-col>
        <b-col class="box">
          <ve-line 
            :data="lineData"
            :settings="lineChartSettings"
          ></ve-line>
        </b-col>
      </b-row>

      <b-row>
        <b-col class="report_drop_down box">
          <b-dropdown size="sm" split :text="dropdownName" class="m-2">
             <b-dropdown-item v-on:click="changeHeatMap(g.id, g.name)" v-for="(g,index) in groupList" :key="index">
               {{g.name}}
             </b-dropdown-item>
          </b-dropdown>
          <ve-heatmap :data="heatData"></ve-heatmap>
        </b-col>
      </b-row>

      <br>
    </b-container>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import tool from '@/lib/tool'

export default {
  name: 'Report',

  data: function () {
    var self = this
    return {
      pieChartSettings: {
        dataType: function (v) {
          return v + ' Min'
        }
      },
      
      pieChartEvents: {
        click: function (e) {
          self.changeTaskChart(e.name)
        }
      },

      lineChartSettings: {},

      pieData: {
        columns: ['Name', 'Min'],
        rows: []
      },

      lineData: {
        columns: ['Date', 'Min'],
        rows: []
      },

      heatData: {
        columns: ['Date', 'Name', 'Min'],
        rows: []
      },
      pie: [],
      line: [],
      heatmap: [],
      groupList: [
        {
          name: "全部",
          id: 0,
        },
      ],
      dropdownName: "全部",
    }
  },

  mixins: [tool],

  computed: {
    ...mapState({
      userInfo: state => state.user.info
    }),
  },

  components: {
  },

  mounted: function () {
    if (this.userInfo.isLogin === false) {
      window.location.href = this.BASE
      return
    }

    this.getGroupList()
    this.getHeatMap() 
    this.getReportLine()
    this.getReportPie()
  },

  methods: {
    getGroupList() {
      let self = this
      this.axios.get(this.APIURL + "/api/groups_name")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.groupList = [...self.groupList, ...resp.data.data.groups]
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    getReportLine() {
      let self = this
      this.axios.get(this.APIURL + "/api/report/line")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.line = resp.data.data.report

        self.initLineData()
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    getReportPie() {
      let self = this
      this.axios.get(this.APIURL + "/api/report/pie")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.pie = resp.data.data.report
        self.initPieChart()
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    changeHeatMap(groupID, name) {
      this.dropdownName = name
      this.heatData.rows = []

      if (groupID === 0)  {
        this.getHeatMap()
      } else {
        this.getTaskHeatMap(groupID)
      }
    },

    getHeatMap() {
      let self = this
      this.axios.get(this.APIURL + "/api/report/group_heatmap")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.heatmap = resp.data.data.report

        self.initHeatMap()
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    getTaskHeatMap(groupID) {
      let self = this
      this.axios.get(this.APIURL + "/api/report/task_heatmap?group_id=" + groupID)
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.heatmap = resp.data.data.report

        self.initHeatMap()
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    initHeatMap() {
      this.heatmap.forEach(d => {
        this.heatData.rows.push({
          "Date": d.date,
          "Name": d.name,
          "Min": Math.floor(d.total_spend_time/60),
        })
      })
    },

    initLineData() {
      this.line.forEach(d => {
        this.lineData.rows.push({
          "Date": d.date,
          "Min": Math.floor(d.total_spend_time/60),
        })
      })
    },

    initPieChart() {
      this.pieData.rows = []
      this.pie.forEach(g => {
        this.pieData.rows.push({
          "Name": g.name,
          "Min": Math.floor(g.total_spend_time/60),
        })
      })
    },

    changeTaskChart(name) {
      this.pie.forEach(g => {
        if (g.name ==name) {  
          this.pieData.rows = []
          g.tasks.forEach(t => {
            this.pieData.rows.push({
              "Name": t.name,
              "Min": Math.floor(t.total_spend_time/60),
            })
          })
        }
      })
    },

    paddingzero(num, length) {
      return (Array(length).join("0") + num).slice(-length)
    },
  }
}
</script>

<style scoped lang="scss">
@import '@/assets/css/variables.scss';

* {
  color: $reportFontColor;
}

#report {
  margin: 0 auto;
  width: 1000px;

  .report_drop_down{
    text-align: left;
  }

  .back_btn{
    cursor: pointer;
    text-align: center;
  }

  .box{
    @extend %reportBox;
    padding-top: 5px;
    padding-bottom: 5px;;
  }
}

</style>