
<template>
<el-container>
    <el-container>
        <el-header height="60px"></el-header>
        <el-main height="">
            <!-- Main content -->
            <el-row>
                <el-col :span="10" offset="1">
                    <el-form :model="clusterConfig" status-icon :rules="rules" ref="clusterConfig" label-width="100px" class="demo-clusterConfig">
                      <el-form-item label="集群名称" prop="name">
                        <el-input v-model="clusterConfig.name" autocomplete="off" placeholder="请输入集群名称">
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
                name: '',
                description: ''
            },
            rules: {
              name: [
                { required: true, message: '请输入集群名', trigger: 'blur' },
              ],
              password: [
                { required: true, message: '请输入登陆密码', trigger: 'blur' },
                { min: 1, message: '登陆密码不能为空', trigger: 'blur' }
              ],
              description: [
                { max: 255, message: '内容不得超过255个字', trigger: 'blur'}
              ]
            }
		}
    },
    methods: {
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$message({
                type: 'success',
                message: this.clusterConfig.name + " 配置成功"
            })
            this.$emit("stepChange")
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
    margin-top: 60px;
}
</style>