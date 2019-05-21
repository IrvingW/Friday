<template>
<el-container>
  <el-container>
    <el-header height="">
      <!-- Header content -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/friday/home/upload_image' }">上传镜像</el-breadcrumb-item>
        <el-breadcrumb-item> </el-breadcrumb-item>
      </el-breadcrumb>
    </el-header>
    <el-main height="">
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
                  action="http://localhost:8080/api/image/package/upload"
                  :on-success="changeFileName"
                  :on-error="uploadFileError"
                  :multiple="false"
                  :limit=1>
                  <i class="el-icon-upload"></i>
                  <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                  <div class="el-upload__tip" slot="tip">只能上传tar文件</div>
                </el-upload>
              </el-col>
            </el-row>
            <el-row style="margin-top: 40px">
                <el-col :span="10" :offset="1">
                    <el-form :model="imageInfo" status-icon :rules="rules" ref="imageInfo" label-width="100px">
                      <el-form-item label="仓库名" prop="repo_name">
                        <el-input v-model="imageInfo.repo_name" autocomplete="off" placeholder="请输入镜像仓库名">
                            <i slot="prefix" class="el-icon-s-home"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item label="标签" prop="tag_name">
                        <el-input v-model="imageInfo.tag_name" autocomplete="off" placeholder="请输入镜像标签名">
                            <i slot="prefix" class="el-icon-s-custom"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item class="button-row">
                        <el-button type="success" @click="submitForm('imageInfo')">提交</el-button>
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
                  action="http://localhost:8080/api/image/dockerfile/upload"
                  :on-success="changeFileName"
                  :on-error="uploadFileError"
                  :multiple="false"
                  :limit=1>
                  <i class="el-icon-upload"></i>
                  <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
                  <div class="el-upload__tip" slot="tip">只能上传tar文件</div>
                </el-upload>
              </el-col>
            </el-row>
            <el-row style="margin-top: 40px">
                <el-col :span="10" :offset="1">
                    <el-form :model="imageInfo" status-icon :rules="rules" ref="imageInfo" label-width="100px">
                      <el-form-item label="仓库名" prop="repo_name">
                        <el-input v-model="imageInfo.repo_name" autocomplete="off" placeholder="请输入镜像仓库名">
                            <i slot="prefix" class="el-icon-s-home"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item label="标签" prop="tag_name">
                        <el-input v-model="imageInfo.tag_name" autocomplete="off" placeholder="请输入镜像标签名">
                            <i slot="prefix" class="el-icon-s-custom"></i>
                        </el-input>
                      </el-form-item>
                      <el-form-item class="button-row">
                        <el-button type="success" @click="submitForm('imageInfo')">提交</el-button>
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
        uploadFileName: "",
        imageInfo: {
          repo_name: "",
          tag_name: ""
        },
        rules: {
          repo_name: [
            { required: true, message: '镜像仓库名不能为空', trigger: 'blur' }
          ],
          tag_name: [
            { required: true, message: '镜像标签名不能为空', trigger: 'blur' }
          ],
        }
		  }
    },
    methods: {
      changeFileName(response, file) {
        uploadFileName = response.body.FileName
        console.log(uploadFileName)
      },
      uploadFileError(err, file) {
        this.$message.error('文件上传失败');
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
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
    }
  }
</script>
<style type="text/css" scoped>
</style>