<template>
  <div id="app">
    <div class="app_nav_bar">
      <b-navbar toggleable="lg" type="light">
        <b-navbar-brand href="#"><router-link to="/">Tomato</router-link></b-navbar-brand>

        <b-collapse id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto" type="light">
            <b-nav-item v-if="this.userState.isLogin" href="#"><router-link to="/report">Report</router-link></b-nav-item>
            <b-nav-item-dropdown :text="userState.name" right>
              <!-- 如果還沒登入就顯示Sign In -->
              <b-dropdown-item v-if="!this.userState.isLogin" v-b-modal.modal-1>Sign In</b-dropdown-item>
              <!-- 如果還登入了就顯示Sign Out -->
              <b-dropdown-item v-else v-on:click="logout">Sign Out</b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
        </b-collapse>
      </b-navbar>
    </div>

    <b-modal id="modal-1" title="Login">
      <h1 v-on:click="googleLogin" class="google_login_btn">Google</h1>
    </b-modal>

    <br>

    <router-view/>
  </div>
</template>
<script>
import { mapState } from 'vuex'
import tool from '@/lib/tool'

export default {
  name: 'App',

  data: function() {
		return {
		};
  },

  mixins: [tool],

  components: {
  },

  computed: {
    ...mapState({
      userState: state => state.user.info
    }),
  },

  mounted: function () {
    this.userInfo()
  },

  methods: {
    userInfo() {
      let self = this
      this.axios.get(this.APIURL + "/api/user/info")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        // 更新user store
        if (resp.data.data.is_login === true) {
          self.$store.dispatch('user/login', {
            id: resp.data.data.user_id,
            name: resp.data.data.user_name,
          })
        }
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    googleLogin() {
      let self = this
      this.axios.get(this.APIURL + "/api/ouath/google/url")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        window.location.replace(resp.data.data.url)
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },

    logout() {
      let self = this
      this.axios.get(this.APIURL + "/api/user/logout")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.Alert(resp.data.msg)
          return
        }

        location.reload();
      })
      .catch(function (error) {
        self.Alert(error)
      });
    },
  },
}
</script>

<style lang="scss">
@import '@/assets/css/variables.scss';

* {
  color: $fontColor;
}

.google_login_btn {
  cursor: pointer;
}

html {
  background-color: $mainColor;
}

.popover-header {
  color: black;
}
.popover-body {
  .b-icon {
    fill: black;     
    font-size: 20px;
  }
}

.modal-body {
  background-color: $mainColor;
}
.modal-header {
  background-color: $mainColor;  
  .close{
    color: $fontColor;
  }
}
.modal-footer {
  background-color: $mainColor; 
}

#app {
  background-color: $mainColor;

  .app_nav_bar {
    margin: 0 auto;
    width: 800px;
    border-bottom: 1px solid $fontColor;
  }

  a {
    text-decoration: none;
    color: inherit;
  }

  .dropdown-item {
    color: black;
  }
}
</style>
