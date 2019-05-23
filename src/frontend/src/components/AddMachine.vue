
<template>
<el-container>
    <el-container>
        <el-header height="60px"></el-header>
        <el-main height="">
            <!-- Main content -->
            <el-row>
                <el-col :span="10" :offset="1">
                    <el-form :model="machineKey" status-icon :rules="rules" ref="machineKey" label-width="100px" class="demo-machineKey">
                      <el-form-item label="IP" prop="ip">
                        <el-input v-model="machineKey.ip" autocomplete="off" placeholder="请输入机器IP地址">
                            <i slot="prefix" class="el-icon-s-home"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item label="用户名" prop="user_name">
                        <el-input v-model="machineKey.user_name" autocomplete="off" placeholder="请输入登陆用户名">
                            <i slot="prefix" class="el-icon-s-custom"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item label="密码" prop="password">
                        <el-input type="password" v-model="machineKey.password" autocomplete="off" show-password placeholder="请输入登陆密码">
                            <i slot="prefix" class="el-icon-key"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item label="端口" prop="port">
                        <el-input v-model="machineKey.port" autocomplete="off" placeholder="请输入ssh端口">
                            <i slot="prefix" class="el-icon-s-opportunity"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item class="button-row">
                        <el-button type="success" @click="submitForm('machineKey')">建立连接</el-button>
                        <el-button @click="resetForm('machineKey')">重置</el-button>
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
  name:'add_machine',
  props:{
    message:{
      type:String
    }
  },
	  data(){
		return{
      machineKey: {
          ip: '',
          user_name: '',
          password: '',
          port: 22
      },
      rules: {
        user_name: [
          { required: true, message: '登陆用户名不能为空', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '登陆密码不能为空', trigger: 'blur' }
        ],
        ip: [
          { required: true, message: 'IP不能为空', trigger: 'blur' }
        ],
        port: [
          { required: true, message: '端口不能为空', trigger: 'blur'}
        ]
      },
      addMachineUrl: '/api/cluster/machine/add'
		}
    },
    methods: {
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      connectMaster() {
        this.$http.post(this.addMachineUrl, this.machineKey,
					{emulateJSON:true, withCredentials: true}).then((response) => {
					  if (response.status != 200){
						  this.$message.error({
							  message: '请求发送失败，请联系管理员'
						  })
					  }
					  var body = response.body
					  if (body.Rtn == 0){
              var master_id = body.MasterId
              this.$emit("masterConnected", master_id)
              this.$emit("stepChange")
              this.$message({
                  type: 'success',
                  message: this.machineKey.ip + " 建立连接成功"
              })
					  }else {
						  this.$message.error({
							  message: body.Msg,
						  })
					  }
          })
      },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.connectMaster()
          } else {
              this.$message.error({
                  message: '请正确填写表单'
              })
            return false;
          }
        });
      }
    }
  }
</script>
<style type="text/css" scoped>
.button-row {
    margin-top: 50px
}
</style>