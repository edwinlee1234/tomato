<template>
  <div class="taks_list">
    <b-modal ref="select_task_modal" hide-footer title="Tasks list">
      <b-card no-body v-for="(group,index) in groups" :key="index" class="mb-1">
        <b-card-header header-tag="header" class="p-1" role="tab">
          <b-button class="group_list" block v-b-toggle="group.name" variant="light">
            {{ group.name }}
          </b-button>
          <span class="group_btn">
            <span><b-icon v-on:click="showEditGroupModal(group)" icon="pencil-square"></b-icon></span>
            <span><b-icon v-on:click="deleteConfirm(group.id)" icon="trash"></b-icon></span>
          </span>  
        </b-card-header>
        <b-collapse :id="group.name" role="tabpanel">
          <b-card-body>
            <b-card-text  v-for="(task,i) in group.tasks" :key="i">
              <span class="task_name" v-on:click="selectTask(task.id, task.name, group.name)">{{ task.name }}</span>
              <span class="task_btn">
                <span><b-icon v-on:click="showEditTaskModal(task)" icon="pencil-square"></b-icon></span>
                <span><b-icon v-on:click="taskDone(task.id)" icon="check2-square"></b-icon></span>
                <span><b-icon v-on:click="deleteConfirm(task.id)" icon="trash"></b-icon></span>
              </span>
            </b-card-text>

            <b-card-text>
              <b-icon class="add_btn" v-on:click="showAddTaskModal(group.id)" id="add_task_btn" icon="plus"></b-icon> 
            </b-card-text>
          </b-card-body>
        </b-collapse>
      </b-card> 

      <b-card no-body class="mb-1">
        <b-card-header header-tag="header" class="p-1" role="tab">
          <b-button v-on:click="showAddGroupModal" block variant="light">
            <b-icon class="add_btn" id="add_task_btn" icon="plus"></b-icon> 
          </b-button>
        </b-card-header>
      </b-card> 
    </b-modal>

    <b-modal @hidden="resetAddTaskModal" ref="add_task_modal" hide-footer title="Add Task">
      <b-form>
        <b-form-group id="task-input-group-1" label="Task Name:" label-for="task-input-1">
          <b-form-input
            id="task-input-1"
            v-model="taskForm.name"
            required
            placeholder="Enter Task name"
          ></b-form-input>
        </b-form-group>

        <b-button variant="primary" v-on:click="addTaskSubmit">Submit</b-button>
      </b-form>
    </b-modal>

    <b-modal @hidden="resetAddTaskModal" ref="add_group_modal" hide-footer title="Add Group">
      <b-form>
        <b-form-group id="group-input-group-1" label="Group Name:" label-for="group-input-1">
          <b-form-input
            id="group-input-1"
            v-model="taskForm.name"
            required
            placeholder="Enter Group name"
          ></b-form-input>
        </b-form-group>

        <b-button variant="primary" v-on:click="addGroupSubmit">Submit</b-button>
      </b-form>
    </b-modal>

    <b-modal ref="edit_task_modal" hide-footer title="Edit Task">
      <b-form>
        <b-form-group id="task-input-group-1" label="Task Name:" label-for="task-input-1">
          <b-form-input
            id="task-input-1"
            v-model="editForm.name"
            required
            placeholder="Enter Task name"
          ></b-form-input>
        </b-form-group>

        <b-button variant="primary" v-on:click="editTaskSubmit">Submit</b-button>
      </b-form>
    </b-modal>

    <b-modal ref="edit_group_modal" hide-footer title="Edit Group">
      <b-form>
        <b-form-group id="group-input-group-1" label="Group Name:" label-for="group-input-1">
          <b-form-input
            id="group-input-1"
            v-model="editForm.name"
            required
            placeholder="Enter Group name"
          ></b-form-input>
        </b-form-group>

        <b-button variant="primary" v-on:click="editGroupSubmit">Submit</b-button>
      </b-form>
    </b-modal>
  </div>
</template>

<script>
import tool from '@/lib/tool'

export default {
  name: 'TaskList',

  mixins: [tool],

  data: function() {
	return {
      tasksList: [],
      editTaskID: 0,
      taskForm: {
        parent_id: 0,
        name: "",
      },
      editForm: {
        id: 0,
        name: "",
      },
      groups: [],
	};
  },

  methods: {
    popupClose() {
      this.$refs.popover.$emit('close')
    },

    showAddTaskModal(parentID) {
      this.taskForm.parent_id = parentID
      this.$refs['add_task_modal'].show()
    },

    showAddGroupModal() {
      this.$refs['add_group_modal'].show()
    },

    closeAddTaskModal() {
      this.$refs['add_task_modal'].hide()
    },

    closeAddGroupModal() {
      this.$refs['add_group_modal'].hide()
    },

    resetAddTaskModal() {
      this.taskForm.name = ""
    },

    showEditTaskModal(task) {
      this.editForm.id = task.id
      this.editForm.name = task.name
      this.$refs['edit_task_modal'].show()
    },

    showEditGroupModal(task) {
      this.editForm.id = task.id
      this.editForm.name = task.name
      this.$refs['edit_group_modal'].show()
    },

    closeEditTaskModal() {
      this.$refs['edit_task_modal'].hide()
    },

    closeEditGroupModal() {
      this.$refs['edit_group_modal'].hide()
    },

    showSelectTaskListModal() {
      if (this.groups.length <= 0) {
        this.fetchGroups()
      }
      this.$refs['select_task_modal'].show()
    },

    closeTaskListModal() {
      this.$refs['select_task_modal'].hide()
    },

    selectTask(id, name, groupName) {
      this.$emit("input",{
        id: id,
        name: name,
        groupName: groupName,
      })

      this.closeTaskListModal()
    },

    taskDone(id) {
      let self = this
      this.axios.put(this.APIURL + "/api/task/done", {
        id: id,
      })
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.fetchGroups()
      })
      .catch(function (error) {
        self.Alert(error)
      })
    },

    addTaskSubmit() {
      let self = this
      if (self.taskForm.name == "") {
          self.Alert("Can't be empty")
          return
      }
      this.axios.post(this.APIURL + "/api/task", {
        parent_id: self.taskForm.parent_id,
        name: self.taskForm.name,
      })
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.fetchGroups()
        self.closeAddTaskModal()
      })
      .catch(function (error) {
        self.Alert(error)
      })
    },

    addGroupSubmit() {
      let self = this
      if (self.taskForm.name == "") {
        self.Alert("Can't be empty")
        return
      }
      this.axios.post(this.APIURL + "/api/task", {
        parent_id: 0,
        name: self.taskForm.name,
      })
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.fetchGroups()
        self.closeAddGroupModal()
      })
      .catch(function (error) {
        self.Alert(error)
      })
    },

    editTaskSubmit() {
      let self = this
      if (self.editForm.name == "") {
        self.Alert("Can't be empty")
        return
      }
      this.axios.put(this.APIURL + "/api/task", {
        id: self.editForm.id,
        name: self.editForm.name,
      })
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.fetchGroups()
        self.closeEditTaskModal()
      })
      .catch(function (error) {
        self.Alert(error)
      })
    },

    editGroupSubmit() {
      let self = this
      if (self.editForm.name == "") {
        self.Alert("Can't be empty")
        return
      }
      this.axios.put(this.APIURL + "/api/task", {
        id: self.editForm.id,
        name: self.editForm.name,
      })
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.fetchGroups()
        self.closeEditGroupModal()
      })
      .catch(function (error) {
        self.Alert(error)
      })
    },

    fetchGroups() {
      let self = this
      this.axios.get(this.APIURL + "/api/groups")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        if (Array.isArray(resp.data.data.groups)) {
          self.groups = resp.data.data.groups
        } else {
          self.groups = []
        }
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    delete(id) {
      let self = this
      this.axios.delete(this.APIURL + "/api/task/" + id)
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        self.fetchGroups()
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    deleteConfirm(id) {
      let self = this
      this.$bvModal.msgBoxConfirm('Please confirm that you want to delete this.', {
        title: 'Please Confirm',
        size: 'sm',
        buttonSize: 'sm',
        okVariant: 'danger',
        okTitle: 'YES',
        cancelTitle: 'NO',
        footerClass: 'p-2',
        hideHeaderClose: false,
        centered: true
      })
        .then(value => {
          if (value === true) {
            self.delete(id)
          }
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

.task_name {
  cursor: pointer;    
}

.group_list {
  width: 80%;
  float: left;
}

.group_btn {
  position: relative;
  float:right;
  margin-top: 6px;
  span {
    padding: 5px;
  }
}

.task_btn {
  position: relative;
  float:right;
  span {
    padding: 5px;
  }
}

</style>