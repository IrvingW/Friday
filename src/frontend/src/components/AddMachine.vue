
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
            }
		}
    },
    methods: {
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      connectMaster() {
          this.$message({
              type: 'success',
              message: this.machineKey.ip + " 建立连接成功"
          })
          var master_id = 2
          this.$emit("stepChange")
          this.$emit("masterConnected", master_id)
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
      },
      changeStep() {
        this.$emit("stepChange")
      }
    }
  }
</script>
<style type="text/css" scoped>
.button-row {
    margin-top: 50px
}
</style>