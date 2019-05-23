
<template>
<el-container>
    <el-container>
        <el-header height="60px"></el-header>
        <el-main height="">
            <!-- Main content -->
            <el-row>
                <el-col :span="10" :offset="1">
                    <el-form :model="clusterConfig" status-icon :rules="rules" ref="clusterConfig" label-width="100px" class="demo-clusterConfig">
                      <el-form-item label="集群名称" prop="name">
                        <el-input v-model="clusterConfig.name" autocomplete="off" placeholder="请输入集群名称">
                        </el-input>
                      </el-form-item>
                      <el-form-item label="路径" prop="path">
                        <el-input v-model="clusterConfig.path" autocomplete="off" placeholder="请输入kubeConfig文件路径">
                        </el-input>
                      </el-form-item>
                      <el-form-item label="端口" prop="port">
                        <el-input v-model="clusterConfig.port" autocomplete="off" placeholder="请输入kubernetes服务端口">
                        </el-input>
                      </el-form-item>
                      <el-form-item label="描述" prop="description">
                          <el-input
                            type="textarea"
                            :autosize="{ minRows: 4, maxRows: 8}"
                            placeholder="请输入对该集群的描述"
                            v-model="clusterConfig.description">
                          </el-input>
                      </el-form-item>
                      <el-form-item class="button-row">
                        <el-button type="success" @click="submitForm('clusterConfig')">提交</el-button>
                        <el-button @click="resetForm('clusterConfig')">重置</el-button>
                      </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
        </el-main>
    </el-container>
</el-container>
</template>
<script>
  export default {
    name:'config_cluster',
	data(){
		return{
      clusterConfig: {
          master_id: -1,
          name: '',
          path: '',
          description: '',
          port: 6443
      },
      rules: {
        name: [
          { required: true, message: '请输入集群名', trigger: 'blur' },
        ],
        path: [
          { required: true, message: '请输入kubeConfig文件路径', trigger: 'blur' },
        ],
        port: [
          { required: true, message: '请输入kubernetes服务端口', trigger: 'blur' },
        ]
      },
      createClusterUrl: '/api/cluster/add'
    }
  },
  methods: {
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
        this.$http.post(this.createClusterUrl, this.clusterConfig,
					{emulateJSON:true, withCredentials: true}).then((response) => {
					  if (response.status != 200){
						  this.$message.error({
							  message: '请求发送失败，请联系管理员'
						  })
					  }
					  var body = response.body
					  if (body.Rtn == 0){

              this.$emit("stepChange")
              this.$message({
                  type: 'success',
                  message: this.clusterConfig.name + " 集群配置成功"
              })
					  }else {
						  this.$message.error({
							  message: body.Msg,
						  })
					  }
          })
        } else {
            this.$message.error({
                message: '请正确填写表单'
            })
          return false;
        }
      });
    }
  },
  created: function () {
    this.clusterConfig.master_id = this.$route.query.master_id
  }
  }
</script>
<style type="text/css" scoped>
.button-row {
    margin-top: 60px;
}
</style>