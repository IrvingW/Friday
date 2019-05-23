<template>
<el-container>
    <el-container>
        <el-header height="">
            <!-- Header content -->
            <el-breadcrumb separator-class="el-icon-arrow-right">
              <el-breadcrumb-item :to="{ path: '/friday/home/add_cluster' }">添加集群</el-breadcrumb-item>
              <el-breadcrumb-item> </el-breadcrumb-item>
            </el-breadcrumb>
        </el-header>
        <el-main height="">
            <!-- Main content -->
            <el-card :body-style="{ padding: '20px' }" shadow="hover">
                <el-row>
                    <el-col :span="18">
                        <el-steps align-center :active="stepActive" finish-status="success">
                          <el-step title="步骤1" description="与Master节点建立连接"></el-step>
                          <el-step title="步骤2" description="配置集群属性"></el-step>
                          <el-step title="步骤4" description="完成"></el-step>
                        </el-steps>
                    </el-col>
                </el-row>
                <el-row>
                    <el-col :span="24">
                        <router-view @stepChange="changeStep()" @masterConnected="setMasterId($event)" v-bind:message="message"></router-view>
                    </el-col>
                </el-row>
            </el-card>
        </el-main>
    </el-container>
</el-container>

</template>
<script>
  export default {
	name:'add_cluster',
	  data(){
		return{
            stepActive: 0,
            master_id: -1,
            router_list: [
                '/friday/home/add_cluster/master',
                '/friday/home/add_cluster/cluster_config',
                '/friday/home/add_cluster/done'
            ],
            message: "asdd"
		}
      },
      methods: {
          changeStep(){
              if (this.stepActive++ === 3){
                  this.stepActive = 0
              }else {
                  if (this.stepActive === 2){
                      this.message = "集群添加成功"
                  }
                  this.$router.push({path: this.router_list[this.stepActive], query: {master_id: this.master_id}})
              }
          },
          setMasterId(master_id){
              this.master_id = master_id
          }
      },
      created: function() {
          this.stepActive = 0
          this.master_id = -1
      }
  }
</script>
<style type="text/css" scoped>
</style>