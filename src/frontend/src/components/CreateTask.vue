<template>
<el-container>
  <el-container>
    <el-header height="">
      <!-- Header content -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/friday/home/create_task' }">创建任务</el-breadcrumb-item>
        <el-breadcrumb-item> </el-breadcrumb-item>
      </el-breadcrumb>
    </el-header>
    <el-main height="">
      <!-- Main content -->
      <el-card :body-style="{ padding: '20px' }" shadow="hover">
        <el-form :model="taskConfig" status-icon :rules="rules" ref="taskConfig" label-width="100px">
          <el-form-item label="任务名称" prop="name">
            <el-input v-model="taskConfig.name" autocomplete="off" placeholder="请输入任务名称">
            </el-input>
          </el-form-item>
          <el-form-item label="任务描述" prop="description">
              <el-input
                type="textarea"
                :autosize="{ minRows: 2, maxRows: 4}"
                placeholder="请输入对该任务的描述"
                v-model="taskConfig.description">
              </el-input>
          </el-form-item>
          <el-form-item label="类名" prop="class_name">
            <el-input v-model="taskConfig.class_name" autocomplete="off" placeholder="请输入类名">
            </el-input>
          </el-form-item>
          <el-form-item label="jar包路径" prop="jar_path">
            <el-input v-model="taskConfig.jar_path" autocomplete="off" placeholder="请输入镜像内的jar包路径">
            </el-input>
          </el-form-item>
          <el-form-item label="调度策略">
              <el-select v-model="taskConfig.scheduler" filterable placeholder="请选择调度策略">
                <el-option
                  v-for="item in scheduler_options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value">
                </el-option>
              </el-select>
          </el-form-item>
          <el-form-item class="button-row">
            <el-button type="success" @click="submitForm('taskConfig')">创建任务</el-button>
            <el-button @click="resetForm('taskConfig')">重置</el-button>
          </el-form-item>
        </el-form>
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
        taskConfig: {
          name: "",
          image_id: 0,
          description: "",
          scheduler: "",
          class_name: "",
          jar_path: ""
        },
        scheduler_options: [
          {value: "roundrobin", label: "轮询调度", key: "roundrobin"},
          {value: "random", label: "随机调度", key: "random"},
          {value: "prefer", label: "偏好调度", key: "prefer"},
          {value: "direct", label: "指定调度", key: "direct"},
          {value: "load", label: "负载均衡", key: "load"},
        ],
        rules: {
          name: [
            { required: true, message: '请输入任务名', trigger: 'blur' },
          ],
          class_name: [
            { required: true, message: '请输入类名', trigger: 'blur' },
          ],
          jar_path: [
            { required: true, message: '请输入镜像内jar包路径', trigger: 'blur' },
          ],
          description: [
            { required: true, message: '请输入对任务的描述', trigger: 'blur' },
          ],
        },
        createTaskUrl: '/api/task/create'
		  }
    },
    methods: {
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
          this.$http.post(this.createTaskUrl, this.taskConfig,
					  {emulateJSON:true, withCredentials: true}).then((response) => {
					    if (response.status != 200){
						    this.$message.error({
							    message: '请求发送失败，请联系管理员'
						    })
					    }
					    var body = response.body
					    if (body.Rtn == 0){
                this.$message({
                    type: 'success',
                    message: this.taskConfig.name + " 任务创建成功"
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
    created: function (params) {
      this.taskConfig.image_id = this.$route.query.image_id
    }
  }
</script>
<style type="text/css" scoped>
</style>