<template>
  <div id="app">
    <div>
      <b-navbar toggleable="lg" type="light">
        <b-navbar-brand href="#">Tomato</b-navbar-brand>
        <b-navbar-nav>
          <b-nav-item><router-link to="/">Home</router-link></b-nav-item>
          <b-nav-item v-if="this.userState.isLogin" href="#"><router-link to="/report">Report</router-link></b-nav-item>
        </b-navbar-nav>

        <b-collapse id="nav-collapse" is-nav>
          <b-navbar-nav class="ml-auto" type="light">
            <b-nav-item-dropdown right>
              <template v-slot:button-content>
                <em>{{userState.name}}</em> <!--顯示會員名稱-->
              </template>
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

export default {
  name: 'App',

  data: function() {
		return {
		};
  },

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
          self.$bvToast.toast(resp.data.msg, {
            title: "Alert",
            variant: "danger",
            solid: true
          })
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
          self.$bvToast.toast(error, {
            title: "Alert",
            variant: "danger",
            solid: true
          })
      });
    },

    googleLogin() {
      let self = this
      this.axios.get(this.APIURL + "/api/ouath/google/url")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.$bvToast.toast(resp.data.msg, {
            title: "Alert",
            variant: "danger",
            solid: true
          })
          return
        }

        window.location.replace(resp.data.data.url)
      })
      .catch(function (error) {
          self.$bvToast.toast(error, {
            title: "Alert",
            variant: "danger",
            solid: true
          })
      });
    },

    logout() {
      let self = this
      this.axios.get(this.APIURL + "/api/user/logout")
      .then(function (resp) {
        if (resp.data.result !== true) {
          self.$bvToast.toast(resp.data.msg, {
            title: "Alert",
            variant: "danger",
            solid: true
          })
          return
        }

        location.reload();
      })
      .catch(function (error) {
         self.$bvToast.toast(error, {
            title: "Alert",
            variant: "danger",
            solid: true
          })
      });
    },
  },
}
</script>

<style lang="scss">
.google_login_btn {
  cursor: pointer;
}
</style>
