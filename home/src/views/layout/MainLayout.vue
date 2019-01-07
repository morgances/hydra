<template>
  <div class="main">
    <div class="content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MainLayout'
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.main {
  width: 100%;
  height: 100%;
}

.content {
  margin-top: 8px;
}

.link {
  text-decoration: none;
  color: black;
}
</style>
