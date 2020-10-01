export default {
  methods: {
    Alert: function (msg) {
      this.$bvToast.toast(msg, {
        title: "Alert",
        variant: "danger",
        solid: true
      })
    }
  }
};