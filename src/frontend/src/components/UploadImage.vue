<template>
<el-container>
  <el-container>
    <el-header>
      <!-- Header content -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/friday/home/repo_list' }">镜像仓库</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/friday/home/repo_view' }">镜像信息</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: '/friday/home/upload_image' }">上传镜像</el-breadcrumb-item>
        <el-breadcrumb-item> </el-breadcrumb-item>
      </el-breadcrumb>
    </el-header>
    <el-main>
      <!-- Main content -->
      <el-card :body-style="{ padding: '30px' }" shadow="hover">
        <el-tabs active-name="package">
          <el-tab-pane label="上传打包镜像" name="package">
            <el-row :gutter="10" style="margin-top: 10px">
              <el-col :span="2" :offset="1">
                <span>上传文件:</span>
              </el-col>
              <el-col :span="12">
                <el-upload
                  drag
                  action="/api/image/package/upload"
                  :on-success="changeFileName"
                  :on-error="uploadFileError"
                  :multiple="false"
                  :limit=1>
                  <i class="el-icon-upload"></i>
                  <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                  <div class="el-upload__tip" slot="tip">只能上传.tar文件</div>
                </el-upload>
              </el-col>
            </el-row>
            <el-row style="margin-top: 40px">
                <el-col :span="10" :offset="1">
                    <el-form :model="imageInfo" status-icon :rules="rules" ref="imageInfo" label-width="100px">
                      <el-form-item label="标签" prop="tag_name">
                        <el-input v-model="imageInfo.tag_name" autocomplete="off" placeholder="请输入镜像标签名">
                            <i slot="prefix" class="el-icon-s-custom"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item label="描述" prop="description">
                        <el-input v-model="imageInfo.description" autocomplete="off" placeholder="请输入镜像简介">
                            <i slot="prefix" class="el-icon-tickets"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item class="button-row">
                        <el-button type="success" @click="submitPackageForm('imageInfo')">提交</el-button>
                        <el-button @click="resetForm('imageInfo')">重置</el-button>
                      </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
          </el-tab-pane>
          <el-tab-pane label="上传Dockerfile" name="file">
            <el-row :gutter="10" style="margin-top: 10px">
              <el-col :span="2" :offset="1">
                <span>上传文件:</span>
              </el-col>
              <el-col :span="12">
                <el-upload
                  drag
                  action="/api/image/dockerfile/upload"
                  :on-success="changeFileName"
                  :on-error="uploadFileError"
                  :multiple="false"
                  :limit=1>
                  <i class="el-icon-upload"></i>
                  <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                  <div class="el-upload__tip" slot="tip">只能上传.zip压缩文件</div>
                </el-upload>
              </el-col>
            </el-row>
            <el-row style="margin-top: 40px">
                <el-col :span="10" :offset="1">
                    <el-form :model="imageInfo" status-icon :rules="rules" ref="imageInfo" label-width="100px">
                      <el-form-item label="标签" prop="tag_name">
                        <el-input v-model="imageInfo.tag_name" autocomplete="off" placeholder="请输入镜像标签名">
                            <i slot="prefix" class="el-icon-s-custom"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item label="描述" prop="description">
                        <el-input v-model="imageInfo.description" autocomplete="off" placeholder="请输入镜像简介">
                            <i slot="prefix" class="el-icon-tickets"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item class="button-row">
                        <el-button type="success" @click="submitFolderForm('imageInfo')">提交</el-button>
                        <el-button @click="resetForm('imageInfo')">重置</el-button>
                      </el-form-item>
                    </el-form>
                </el-col>
            </el-row>
          </el-tab-pane>
        </el-tabs>
      </el-card>
    </el-main>
  </el-container>
</el-container>

</template>
<script>
  export default {
	name:'upload_image',
	  data(){
		  return{
        imageInfo: {
          repo_id: 0,
          description: "",
          tag_name: "",
          path: ""
        },
        rules: {
          tag_name: [
            { required: true, message: '镜像标签名不能为空', trigger: 'blur' }
          ],
        },
        addImagePackageUrl: "/api/image/add/package",
        addImageDockerfileUrl: "/api/image/add/dockerfile"
		  }
    },
    created: function () {
      this.imageInfo.repo_id = this.$route.query.repo_id
    },
    methods: {
      changeFileName(response, file) {
				if (response.Rtn == 0){
          this.imageInfo.path = response.FilePath
					this.$message({
							message: "文件上传成功",
               type: 'success'
					});
				}else{
					this.$message.error({
               message: response.Msg
						})
				}
      },
      uploadFileError(err, file) {
        this.$message.error('文件上传失败');
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      },
      submitPackageForm(formName) {
        if (this.imageInfo.path == "") {
          this.$message.error({message: "请先上传文件"})
          return
        }
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$http.post(this.addImagePackageUrl,
							this.imageInfo,
							{emulateJSON:true, withCredentials: true}
            ).then((response) => {
							if (response.status != 200){
								this.$message.error({
									message: '请求发送失败，请联系管理员'
								})
							}
							var body = response.body
							if (body.Rtn == 0){
                this.$message({
                  type: "success",
                  message: "镜像上传成功"
                })
							}else {
								this.$message.error({
                  message: body.Msg
                });
              }
            })
          } else {
              this.$message.error({
                  message: '请正确填写表单'
              })
            return false;
          }
        });
      },
      submitFolderForm(formName) {
        if (this.imageInfo.path == "") {
          this.$message.error({message: "请先上传文件"})
          return
        }
        this.$refs[formName].validate((valid) => {
          if (valid) {
            this.$http.post(this.addImageDockerfileUrl,
							this.imageInfo,
							{emulateJSON:true, withCredentials: true}
            ).then((response) => {
							if (response.status != 200){
								this.$message.error({
									message: '请求发送失败，请联系管理员'
								})
							}
							var body = response.body
							if (body.Rtn == 0){
                this.$message({
                  type: "success",
                  message: "镜像上传成功"
                })
							}else {
								this.$message.error({
                  message: body.Msg
                });
              }
            })
          } else {
              this.$message.error({
                  message: '请正确填写表单'
              })
            return false;
          }
        });
      },
    }
  }
</script>
<style type="text/css" scoped>
</style>