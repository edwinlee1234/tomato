(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-33f6892e"],{"0585":function(t,e,i){},"1a3c":function(t,e,i){"use strict";var a=i("85e7"),r=i.n(a);r.a},"531a":function(t,e,i){t.exports=i("bb56").default},5899:function(t,e){t.exports="\t\n\v\f\r                　\u2028\u2029\ufeff"},"58a8":function(t,e,i){var a=i("1d80"),r=i("5899"),s="["+r+"]",n=RegExp("^"+s+s+"*"),o=RegExp(s+s+"*$"),c=function(t){return function(e){var i=String(a(e));return 1&t&&(i=i.replace(n,"")),2&t&&(i=i.replace(o,"")),i}};t.exports={start:c(1),end:c(2),trim:c(3)}},"6b97":function(t,e,i){},7156:function(t,e,i){var a=i("861d"),r=i("d2bb");t.exports=function(t,e,i){var s,n;return r&&"function"==typeof(s=e.constructor)&&s!==i&&a(n=s.prototype)&&n!==i.prototype&&r(t,n),t}},"7d89":function(t,e,i){"use strict";var a=i("0585"),r=i.n(a);r.a},"85bf":function(t,e,i){"use strict";var a=i("6b97"),r=i.n(a);r.a},"85e7":function(t,e,i){},"99af":function(t,e,i){"use strict";var a=i("23e7"),r=i("d039"),s=i("e8b5"),n=i("861d"),o=i("7b0b"),c=i("50c4"),u=i("8418"),l=i("65f0"),d=i("1dde"),h=i("b622"),f=i("2d00"),m=h("isConcatSpreadable"),p=9007199254740991,b="Maximum allowed index exceeded",k=f>=51||!r((function(){var t=[];return t[m]=!1,t.concat()[0]!==t})),g=d("concat"),v=function(t){if(!n(t))return!1;var e=t[m];return void 0!==e?!!e:s(t)},_=!k||!g;a({target:"Array",proto:!0,forced:_},{concat:function(t){var e,i,a,r,s,n=o(this),d=l(n,0),h=0;for(e=-1,a=arguments.length;e<a;e++)if(s=-1===e?n:arguments[e],v(s)){if(r=c(s.length),h+r>p)throw TypeError(b);for(i=0;i<r;i++,h++)i in s&&u(d,h,s[i])}else{if(h>=p)throw TypeError(b);u(d,h++,s)}return d.length=h,d}})},a15b:function(t,e,i){"use strict";var a=i("23e7"),r=i("44ad"),s=i("fc6a"),n=i("a640"),o=[].join,c=r!=Object,u=n("join",",");a({target:"Array",proto:!0,forced:c||!u},{join:function(t){return o.call(s(this),void 0===t?",":t)}})},a9e3:function(t,e,i){"use strict";var a=i("83ab"),r=i("da84"),s=i("94ca"),n=i("6eeb"),o=i("5135"),c=i("c6b6"),u=i("7156"),l=i("c04e"),d=i("d039"),h=i("7c73"),f=i("241c").f,m=i("06cf").f,p=i("9bf2").f,b=i("58a8").trim,k="Number",g=r[k],v=g.prototype,_=c(h(v))==k,S=function(t){var e,i,a,r,s,n,o,c,u=l(t,!1);if("string"==typeof u&&u.length>2)if(u=b(u),e=u.charCodeAt(0),43===e||45===e){if(i=u.charCodeAt(2),88===i||120===i)return NaN}else if(48===e){switch(u.charCodeAt(1)){case 66:case 98:a=2,r=49;break;case 79:case 111:a=8,r=55;break;default:return+u}for(s=u.slice(2),n=s.length,o=0;o<n;o++)if(c=s.charCodeAt(o),c<48||c>r)return NaN;return parseInt(s,a)}return+u};if(s(k,!g(" 0o1")||!g("0b1")||g("+0x1"))){for(var A,y=function(t){var e=arguments.length<1?0:t,i=this;return i instanceof y&&(_?d((function(){v.valueOf.call(i)})):c(i)!=k)?u(new g(S(e)),i,y):S(e)},T=a?f(g):"MAX_VALUE,MIN_VALUE,NaN,NEGATIVE_INFINITY,POSITIVE_INFINITY,EPSILON,isFinite,isInteger,isNaN,isSafeInteger,MAX_SAFE_INTEGER,MIN_SAFE_INTEGER,parseFloat,parseInt,isInteger".split(","),C=0;T.length>C;C++)o(g,A=T[C])&&!o(y,A)&&p(y,A,m(g,A));y.prototype=v,v.constructor=y,n(r,k,y)}},bb51:function(t,e,i){"use strict";i.r(e);var a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",{staticClass:"home"},[i("div",{attrs:{id:"time_block"}},[i("b-container",{attrs:{id:"select_timer"}},[i("b-row",[i("b-col",{attrs:{id:"long_break"},on:{click:t.longBreak}},[i("b-button",{attrs:{variant:"outline-light"}},[t._v("Long Break")])],1),i("b-col",{attrs:{id:"shot_break"},on:{click:t.shotBreak}},[i("b-button",{attrs:{variant:"outline-light"}},[t._v("Short Break")])],1),i("b-col",{attrs:{id:"custom"}},[i("b-button",{attrs:{variant:"outline-light"}},[t._v("Custom")])],1)],1)],1),i("div",{attrs:{id:"time"}},[i("radial-progress-bar",{attrs:{diameter:200,"completed-steps":t.runTime,"total-steps":t.startTimeVal}},[i("p",[t._v(t._s(t.paddingzero(Math.floor(t.timeVal/60),2))+":"+t._s(t.paddingzero(t.timeVal%60,2)))])])],1),t.isStart?t._e():i("b-button",{staticClass:"btn",attrs:{variant:"outline-light"},on:{click:t.start}},[t._v("Start")]),t.isStart?i("b-button",{staticClass:"btn",attrs:{variant:"danger"},on:{click:t.stop}},[t._v("Stop")]):t._e(),t.isStart?t._e():i("b-button",{staticClass:"btn",attrs:{variant:"outline-light"},on:{click:t.reset}},[t._v("Reset")]),i("b-popover",{ref:"popover",attrs:{target:"custom",title:"Minute"}},[i("b-form-input",{attrs:{placeholder:"Minute"},model:{value:t.customMinute,callback:function(e){t.customMinute=e},expression:"customMinute"}}),i("br"),i("b-icon",{staticClass:"h5 mb-2",attrs:{icon:"check"},on:{click:t.custom}},[t._v("OK")]),i("b-icon",{staticClass:"h5 mb-2",attrs:{icon:"x"},on:{click:t.popupClose}},[t._v("X")])],1)],1),i("br"),t.userInfo.isLogin?i("div",{attrs:{id:"task_list"}},[t.taskItem?i("div",[i("h4",[t._v("Task")]),i("hr"),i("b-list-group",[i("b-list-group-item",[t._v(t._s(t.taskItem.groupName)+" - "+t._s(t.taskItem.name)+" "),i("span",{staticClass:"task_btn"},[i("b-icon",{attrs:{icon:"pencil-square"},on:{click:function(e){return t.showSelectTaskListModal()}}}),i("b-icon",{attrs:{icon:"trash"},on:{click:function(e){return t.clearSelectedTask()}}})],1)])],1)],1):i("div",[i("h4",[t._v("Choose a task")]),i("hr"),i("h4",[i("b-icon",{staticClass:"plus_btn",attrs:{icon:"plus"},on:{click:function(e){return t.showSelectTaskListModal()}}})],1)])]):t._e(),i("TaskList",{ref:"TaskList",on:{input:t.setTaskItem}}),i("br"),t.userInfo.isLogin?i("div",{attrs:{id:"today_record"}},[i("div",[i("h4",[t._v("Today")]),i("hr"),i("b-list-group",t._l(t.todayRecords,(function(e,a){return i("b-list-group-item",{key:a},[t._v(" "+t._s(t.convertTimestampToDate(e.created_timestamp-e.spend_time))+" "+t._s(e.group_name)+" - "+t._s(e.name)+" "+t._s(Math.floor(e.spend_time/60))+" Minutes ")])})),1)],1)]):t._e()],1)},r=[],s=(i("a15b"),i("fb6a"),i("5530")),n=i("2f62"),o=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",{staticClass:"taks_list"},[i("b-modal",{ref:"select_task_modal",attrs:{"hide-footer":"",title:"Tasks list"}},[t._l(t.groups,(function(e,a){return i("b-card",{key:a,staticClass:"mb-1",attrs:{"no-body":""}},[i("b-card-header",{staticClass:"p-1",attrs:{"header-tag":"header",role:"tab"}},[i("b-button",{directives:[{name:"b-toggle",rawName:"v-b-toggle",value:e.name,expression:"group.name"}],staticClass:"group_list",attrs:{block:"",variant:"light"}},[t._v(" "+t._s(e.name)+" ")]),i("span",{staticClass:"group_btn"},[i("span",[i("b-icon",{attrs:{icon:"pencil-square"},on:{click:function(i){return t.showEditGroupModal(e)}}})],1),i("span",[i("b-icon",{attrs:{icon:"trash"},on:{click:function(i){return t.deleteConfirm(e.id)}}})],1)])],1),i("b-collapse",{attrs:{id:e.name,role:"tabpanel"}},[i("b-card-body",[t._l(e.tasks,(function(a,r){return i("b-card-text",{key:r},[i("div",{staticClass:"task_item"},[i("span",{staticClass:"task_name",on:{click:function(i){return t.selectTask(a.id,a.name,e.name)}}},[t._v(t._s(a.name))]),i("span",{staticClass:"task_btn"},[i("span",[i("b-icon",{attrs:{icon:"pencil-square"},on:{click:function(e){return t.showEditTaskModal(a)}}})],1),i("span",[i("b-icon",{attrs:{icon:"check2-square"},on:{click:function(e){return t.taskDone(a.id)}}})],1),i("span",[i("b-icon",{attrs:{icon:"trash"},on:{click:function(e){return t.deleteConfirm(a.id)}}})],1)])])])})),i("b-card-text",[i("b-icon",{staticClass:"add_task_btn",attrs:{icon:"plus"},on:{click:function(i){return t.showAddTaskModal(e.id)}}})],1)],2)],1)],1)})),i("b-card",{staticClass:"mb-1",attrs:{"no-body":""}},[i("b-card-header",{staticClass:"p-1",attrs:{"header-tag":"header",role:"tab"}},[i("b-button",{attrs:{block:"",variant:"light"},on:{click:t.showAddGroupModal}},[i("b-icon",{attrs:{icon:"plus"}})],1)],1)],1)],2),i("b-modal",{ref:"add_task_modal",attrs:{"hide-footer":"",title:"Add Task"},on:{hidden:t.resetAddTaskModal}},[i("b-form",[i("b-form-group",{attrs:{id:"task-input-group-1",label:"Task Name:","label-for":"task-input-1"}},[i("b-form-input",{attrs:{id:"task-input-1",required:"",placeholder:"Enter Task name"},model:{value:t.taskForm.name,callback:function(e){t.$set(t.taskForm,"name",e)},expression:"taskForm.name"}})],1),i("br"),i("b-button",{attrs:{variant:"outline-light"},on:{click:t.addTaskSubmit}},[t._v("Submit")])],1)],1),i("b-modal",{ref:"add_group_modal",attrs:{"hide-footer":"",title:"Add Group"},on:{hidden:t.resetAddTaskModal}},[i("b-form",[i("b-form-group",{attrs:{id:"group-input-group-1",label:"Group Name:","label-for":"group-input-1"}},[i("b-form-input",{attrs:{id:"group-input-1",required:"",placeholder:"Enter Group name"},model:{value:t.taskForm.name,callback:function(e){t.$set(t.taskForm,"name",e)},expression:"taskForm.name"}})],1),i("br"),i("b-button",{attrs:{variant:"outline-light"},on:{click:t.addGroupSubmit}},[t._v("Submit")])],1)],1),i("b-modal",{ref:"edit_task_modal",attrs:{"hide-footer":"",title:"Edit Task"}},[i("b-form",[i("b-form-group",{attrs:{id:"task-input-group-1",label:"Task Name:","label-for":"task-input-1"}},[i("b-form-input",{attrs:{id:"task-input-1",required:"",placeholder:"Enter Task name"},model:{value:t.editForm.name,callback:function(e){t.$set(t.editForm,"name",e)},expression:"editForm.name"}})],1),i("br"),i("b-button",{attrs:{variant:"outline-light"},on:{click:t.editTaskSubmit}},[t._v("Submit")])],1)],1),i("b-modal",{ref:"edit_group_modal",attrs:{"hide-footer":"",title:"Edit Group"}},[i("b-form",[i("b-form-group",{attrs:{id:"group-input-group-1",label:"Group Name:","label-for":"group-input-1"}},[i("b-form-input",{attrs:{id:"group-input-1",required:"",placeholder:"Enter Group name"},model:{value:t.editForm.name,callback:function(e){t.$set(t.editForm,"name",e)},expression:"editForm.name"}})],1),i("br"),i("b-button",{attrs:{variant:"outline-light"},on:{click:t.editGroupSubmit}},[t._v("Submit")])],1)],1)],1)},c=[],u=(i("b0c0"),i("b4eb")),l={name:"TaskList",mixins:[u["a"]],data:function(){return{tasksList:[],editTaskID:0,taskForm:{parent_id:0,name:""},editForm:{id:0,name:""},groups:[]}},methods:{popupClose:function(){this.$refs.popover.$emit("close")},showAddTaskModal:function(t){this.taskForm.parent_id=t,this.$refs["add_task_modal"].show()},showAddGroupModal:function(){this.$refs["add_group_modal"].show()},closeAddTaskModal:function(){this.$refs["add_task_modal"].hide()},closeAddGroupModal:function(){this.$refs["add_group_modal"].hide()},resetAddTaskModal:function(){this.taskForm.name=""},showEditTaskModal:function(t){this.editForm.id=t.id,this.editForm.name=t.name,this.$refs["edit_task_modal"].show()},showEditGroupModal:function(t){this.editForm.id=t.id,this.editForm.name=t.name,this.$refs["edit_group_modal"].show()},closeEditTaskModal:function(){this.$refs["edit_task_modal"].hide()},closeEditGroupModal:function(){this.$refs["edit_group_modal"].hide()},showSelectTaskListModal:function(){this.groups.length<=0&&this.fetchGroups(),this.$refs["select_task_modal"].show()},closeTaskListModal:function(){this.$refs["select_task_modal"].hide()},selectTask:function(t,e,i){this.$emit("input",{id:t,name:e,groupName:i}),this.closeTaskListModal()},taskDone:function(t){var e=this;this.axios.put(this.APIURL+"/api/task/done",{id:t}).then((function(t){!0===t.data.result?e.fetchGroups():e.Alert(t.data.msg)})).catch((function(t){e.Alert(t)}))},addTaskSubmit:function(){var t=this;""!=t.taskForm.name?this.axios.post(this.APIURL+"/api/task",{parent_id:t.taskForm.parent_id,name:t.taskForm.name}).then((function(e){!0===e.data.result?(t.fetchGroups(),t.closeAddTaskModal()):t.Alert(e.data.msg)})).catch((function(e){t.Alert(e)})):t.Alert("Can't be empty")},addGroupSubmit:function(){var t=this;""!=t.taskForm.name?this.axios.post(this.APIURL+"/api/task",{parent_id:0,name:t.taskForm.name}).then((function(e){!0===e.data.result?(t.fetchGroups(),t.closeAddGroupModal()):t.Alert(e.data.msg)})).catch((function(e){t.Alert(e)})):t.Alert("Can't be empty")},editTaskSubmit:function(){var t=this;""!=t.editForm.name?this.axios.put(this.APIURL+"/api/task",{id:t.editForm.id,name:t.editForm.name}).then((function(e){!0===e.data.result?(t.fetchGroups(),t.closeEditTaskModal()):t.Alert(e.data.msg)})).catch((function(e){t.Alert(e)})):t.Alert("Can't be empty")},editGroupSubmit:function(){var t=this;""!=t.editForm.name?this.axios.put(this.APIURL+"/api/task",{id:t.editForm.id,name:t.editForm.name}).then((function(e){!0===e.data.result?(t.fetchGroups(),t.closeEditGroupModal()):t.Alert(e.data.msg)})).catch((function(e){t.Alert(e)})):t.Alert("Can't be empty")},fetchGroups:function(){var t=this;this.axios.get(this.APIURL+"/api/groups").then((function(e){!0===e.data.result?Array.isArray(e.data.data.groups)?t.groups=e.data.data.groups:t.groups=[]:t.Alert(e.data.msg)})).catch((function(e){t.Alert(e)}))},delete:function(t){var e=this;this.axios.delete(this.APIURL+"/api/task/"+t).then((function(t){!0===t.data.result?e.fetchGroups():e.Alert(t.data.msg)})).catch((function(t){e.Alert(t)}))},deleteConfirm:function(t){var e=this;this.$bvModal.msgBoxConfirm("Please confirm that you want to delete this.",{title:"Please Confirm",size:"sm",buttonSize:"sm",okVariant:"danger",okTitle:"YES",cancelTitle:"NO",footerClass:"p-2",hideHeaderClose:!1,centered:!0}).then((function(i){!0===i&&e.delete(t)}))}}},d=l,h=(i("7d89"),i("2877")),f=Object(h["a"])(d,o,c,!1,null,"eed4107a",null),m=f.exports,p=i("531a"),b=i.n(p),k={name:"Home",data:function(){return{timeVal:1500,timeInterval:null,isStart:!1,customMinute:5,startTimeVal:1500,taskItem:null,todayRecords:[],runTime:0}},mixins:[u["a"]],components:{TaskList:m,RadialProgressBar:b.a},computed:Object(s["a"])({},Object(n["b"])({userInfo:function(t){return t.user.info}})),watch:{userInfo:function(){!0===this.userInfo.isLogin&&this.fetchTodayRecords()}},mounted:function(){!1!==this.userInfo.isLogin&&this.fetchTodayRecords()},methods:{start:function(){0!==this.timeVal&&(this.timeInterval=setInterval(this.countDown,1e3),this.isStart=!0)},countDown:function(){this.timeVal--,this.runTime++,0===this.timeVal&&this.finish()},stop:function(){clearInterval(this.timeInterval),this.isStart=!1},reset:function(){this.timeVal=this.startTimeVal,this.runTime=0},shotBreak:function(){this.stop(),this.timeVal=900,this.runTime=0,this.startTimeVal=this.timeVal},longBreak:function(){this.stop(),this.timeVal=1500,this.runTime=0,this.startTimeVal=this.timeVal},custom:function(){this.stop(),this.customMinute>60&&(this.customMinute=60),this.timeVal=60*this.customMinute,this.runTime=0,this.startTimeVal=this.timeVal,this.popupClose()},paddingzero:function(t,e){return(Array(e).join("0")+t).slice(-e)},popupClose:function(){this.$refs.popover.$emit("close")},showSelectTaskListModal:function(){this.$refs.TaskList.showSelectTaskListModal()},setTaskItem:function(t){this.taskItem=t,this.stop(),this.reset()},clearSelectedTask:function(){this.taskItem=null,this.stop(),this.reset()},convertTimestampToDate:function(t){var e=new Date(1e3*t),i=e.getHours(),a="0"+e.getMinutes(),r="0"+e.getSeconds();return i+":"+a.substr(-2)+":"+r.substr(-2)},finish:function(){if(this.stop(),null!==this.taskItem){var t=this;this.axios.post(this.APIURL+"/api/pomo",{id:this.taskItem.id,time:this.startTimeVal}).then((function(e){!0===e.data.result?t.fetchTodayRecords():t.Alert(e.data.msg)})).catch((function(e){t.Alert(e)}))}},fetchTodayRecords:function(){var t=this;this.axios.get(this.APIURL+"/api/day/record").then((function(e){!0===e.data.result?t.todayRecords=e.data.data.records:t.Alert(e.data.msg)})).catch((function(e){t.Alert(e)}))}}},g=k,v=(i("1a3c"),Object(h["a"])(g,a,r,!1,null,"2108659a",null));e["default"]=v.exports},bb56:function(t,e,i){"use strict";i.r(e);var a=function(){var t=this,e=t.$createElement,i=t._self._c||e;return i("div",{staticClass:"radial-progress-container",style:t.containerStyle},[i("div",{staticClass:"radial-progress-inner",style:t.innerCircleStyle},[t._t("default")],2),i("svg",{staticClass:"radial-progress-bar",attrs:{width:t.diameter,height:t.diameter,version:"1.1",xmlns:"http://www.w3.org/2000/svg"}},[i("defs",[i("radialGradient",{attrs:{id:"radial-gradient"+t._uid,fx:t.gradient.fx,fy:t.gradient.fy,cx:t.gradient.cx,cy:t.gradient.cy,r:t.gradient.r}},[i("stop",{attrs:{offset:"30%","stop-color":t.startColor}}),i("stop",{attrs:{offset:"100%","stop-color":t.stopColor}})],1)],1),i("circle",{style:t.strokeStyle,attrs:{r:t.innerCircleRadius,cx:t.radius,cy:t.radius,fill:"transparent",stroke:t.innerStrokeColor,"stroke-dasharray":t.circumference,"stroke-dashoffset":"0","stroke-linecap":t.strokeLinecap}}),i("circle",{style:t.progressStyle,attrs:{transform:"rotate(270, "+t.radius+","+t.radius+")",r:t.innerCircleRadius,cx:t.radius,cy:t.radius,fill:"transparent",stroke:"url(#radial-gradient"+t._uid+")","stroke-dasharray":t.circumference,"stroke-dashoffset":t.circumference,"stroke-linecap":t.strokeLinecap}})])])},r=[],s=(i("99af"),i("a9e3"),{props:{diameter:{type:Number,required:!1,default:200},totalSteps:{type:Number,required:!0,default:10},completedSteps:{type:Number,required:!0,default:0},startColor:{type:String,required:!1,default:"#bbff42"},stopColor:{type:String,required:!1,default:"#429321"},strokeWidth:{type:Number,required:!1,default:10},innerStrokeWidth:{type:Number,required:!1,default:10},strokeLinecap:{type:String,required:!1,default:"round"},animateSpeed:{type:Number,required:!1,default:1e3},innerStrokeColor:{type:String,required:!1,default:"#323232"},fps:{type:Number,required:!1,default:60},timingFunc:{type:String,required:!1,default:"linear"},isClockwise:{type:Boolean,required:!1,default:!0}},data:function(){return{gradient:{fx:.99,fy:.5,cx:.5,cy:.5,r:.65},gradientAnimation:null,currentAngle:0,strokeDashoffset:0}},computed:{radius:function(){return this.diameter/2},circumference:function(){return Math.PI*this.innerCircleDiameter},stepSize:function(){return 0===this.totalSteps?0:100/this.totalSteps},finishedPercentage:function(){return this.stepSize*this.completedSteps},circleSlice:function(){return 2*Math.PI/this.totalSteps},animateSlice:function(){return this.circleSlice/this.totalPoints},innerCircleDiameter:function(){return this.diameter-2*this.innerStrokeWidth},innerCircleRadius:function(){return this.innerCircleDiameter/2},totalPoints:function(){return this.animateSpeed/this.animationIncrements},animationIncrements:function(){return 1e3/this.fps},hasGradient:function(){return this.startColor!==this.stopColor},containerStyle:function(){return{height:"".concat(this.diameter,"px"),width:"".concat(this.diameter,"px")}},progressStyle:function(){return{height:"".concat(this.diameter,"px"),width:"".concat(this.diameter,"px"),strokeWidth:"".concat(this.strokeWidth,"px"),strokeDashoffset:this.strokeDashoffset,transition:"stroke-dashoffset ".concat(this.animateSpeed,"ms ").concat(this.timingFunc)}},strokeStyle:function(){return{height:"".concat(this.diameter,"px"),width:"".concat(this.diameter,"px"),strokeWidth:"".concat(this.innerStrokeWidth,"px")}},innerCircleStyle:function(){return{width:"".concat(this.innerCircleDiameter,"px")}}},methods:{getStopPointsOfCircle:function(t){for(var e=[],i=0;i<t;i++){var a=this.circleSlice*i;e.push(this.getPointOfCircle(a))}return e},getPointOfCircle:function(t){var e=.5,i=e+e*Math.cos(t),a=e+e*Math.sin(t);return{x:i,y:a}},gotoPoint:function(){var t=this.getPointOfCircle(this.currentAngle);t.x&&t.y&&(this.gradient.fx=t.x,this.gradient.fy=t.y)},direction:function(){return this.isClockwise?1:-1},changeProgress:function(t){var e=this,i=t.isAnimate,a=void 0===i||i;if(this.strokeDashoffset=(100-this.finishedPercentage)/100*this.circumference*this.direction(),this.gradientAnimation&&clearInterval(this.gradientAnimation),a){var r=(this.completedSteps-1)*this.circleSlice,s=(this.currentAngle-r)/this.animateSlice,n=Math.abs(s-this.totalPoints)/this.totalPoints,o=s<this.totalPoints;this.gradientAnimation=setInterval((function(){o&&s>=e.totalPoints||!o&&s<e.totalPoints?clearInterval(e.gradientAnimation):(e.currentAngle=r+e.animateSlice*s,e.gotoPoint(),s+=o?n:-n)}),this.animationIncrements)}else this.gotoNextStep()},gotoNextStep:function(){this.currentAngle=this.completedSteps*this.circleSlice,this.gotoPoint()}},watch:{totalSteps:function(){this.changeProgress({isAnimate:!0})},completedSteps:function(){this.changeProgress({isAnimate:!0})},diameter:function(){this.changeProgress({isAnimate:!0})},strokeWidth:function(){this.changeProgress({isAnimate:!0})}},created:function(){this.changeProgress({isAnimate:!1})}}),n=s,o=(i("85bf"),i("2877")),c=Object(o["a"])(n,a,r,!1,null,null,null);e["default"]=c.exports},fb6a:function(t,e,i){"use strict";var a=i("23e7"),r=i("861d"),s=i("e8b5"),n=i("23cb"),o=i("50c4"),c=i("fc6a"),u=i("8418"),l=i("b622"),d=i("1dde"),h=i("ae40"),f=d("slice"),m=h("slice",{ACCESSORS:!0,0:0,1:2}),p=l("species"),b=[].slice,k=Math.max;a({target:"Array",proto:!0,forced:!f||!m},{slice:function(t,e){var i,a,l,d=c(this),h=o(d.length),f=n(t,h),m=n(void 0===e?h:e,h);if(s(d)&&(i=d.constructor,"function"!=typeof i||i!==Array&&!s(i.prototype)?r(i)&&(i=i[p],null===i&&(i=void 0)):i=void 0,i===Array||void 0===i))return b.call(d,f,m);for(a=new(void 0===i?Array:i)(k(m-f,0)),l=0;f<m;f++,l++)f in d&&u(a,l,d[f]);return a.length=l,a}})}}]);
//# sourceMappingURL=chunk-33f6892e.9a0f2027.js.map