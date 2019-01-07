<template>
  <div>
    <el-row  v-for="(o) in 3" :key="o" class="course">
      <el-col :span="24">
        <el-card :body-style="{ padding: '0px' }">
          <img src="@/assets/teaching.svg" class="image">
          <div style="padding: 14px;">
            <span>语文、数学、英语</span>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
export default {
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}

.title {
  font-size: 24px;
}

.image {
  width: 80%;
  margin: auto;
  display: block;
}

.course {
  margin-top: 12px;
}
</style>
