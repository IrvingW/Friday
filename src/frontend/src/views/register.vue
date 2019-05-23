<template>
	<el-container>
    <el-main>
	  <div class="login-wrap">
		  <Background></Background>
		  <div id="login">
			  <h2 class="login-title">用户注册</h2>
			  <el-form label-width="100px" :model="registerForm" :rules="rules" ref="registerForm">
			    <el-form-item label="验证码" prop="token">
			      <el-input v-model="registerForm.token" clearable></el-input>
			    </el-form-item>
			    <el-form-item label="用户名" prop="username">
			      <el-input v-model="registerForm.username" clearable></el-input>
			    </el-form-item>
			    <el-form-item label="密码" prop="password">
			      <el-input v-model="registerForm.password" type="password"></el-input>
			    </el-form-item>
			    <el-form-item label="确认密码" prop="ensure_pwd">
			      <el-input v-model="registerForm.ensure_pwd" type="password"></el-input>
			    </el-form-item>
					<el-form-item>
            <el-button type="primary" @click="submitForm('registerForm')">立即注册</el-button>
            <el-button @click="resetForm('registerForm')">重置</el-button>
          </el-form-item>
			  </el-form>
		  </div>

	  </div>
    </el-main>
	</el-container>
</template>
<script>
	import Background from '@/components/Background'
	export default {
		name: 'Login',
		components:{Background},
		data() {
				var validatePass = (rule, value, callback) => {
             if (this.registerForm.ensure_pwd !== '') {
               this.$refs.registerForm.validateField('ensure_pwd');
             }
             callback();
         };
        var validatePass2 = (rule, value, callback) => {
          if (value !== this.registerForm.password) {
            callback(new Error('两次输入密码不一致!'));
          } else {
            callback();
          }
        };
	      return {
	        registerForm: {
						token: '',
	          username: '',
						password: '',
						ensure_pwd: '',
					},
	        rules:{
						token: [
	        		{required:true,message:'请输入验证码',trigger:'blur'},
	        		{min:1,message:'验证码不能为空',trigger:'blur'}
						],
	        	username:[
	        		{required:true,message:'请输入用户名',trigger:'blur'},
	        		{min:1,message:'用户名不能为空',trigger:'blur'}
	        	],
	        	password:[
	        		{required:true,message:'请输入密码',trigger:'blur'},
							{min:1,message:'密码不能为空',trigger:'blur'},
							{validator: validatePass, trigger: 'blur'}
						],
						ensure_pwd:[
	        		{required:true,message:'请确认密码',trigger:'blur'},
	        		{min:1,message:'密码不能为空',trigger:'blur'},
							{validator: validatePass2, trigger: 'blur'}
						]
					},
					register_url: "/api/user/register"
	      }
		},
		methods: {
			resetForm(formName) {
				this.$refs[formName].resetFields();
			},
			submitForm(formName) {
				this.$refs[formName].validate((valid) => {
					if (valid){
						this.$http.post(
							this.register_url,
							this.registerForm,
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
                  message: '注册成功，请转至登陆界面登陆',
                  type: 'success'
								});
								this.$router.push({path: "/friday/login"})
							}else if (body.Rtn < 0){
								this.$message.error({
                  message: body.Msg
                });
							}else {
								this.$message({
									message: body.Msg,
									type: 'warning'
								})
							}
						})
					}else{
						alert("请正确填写表单之后再提交！")
						resetForm('registerForm')
					}
				})
			}

	  }
	}
</script>



<style scoped>
	.login-wrap >>> label {
		font-size: 13px;
		font-weight: 200;
	}
	#login {
	    background-color: #ffffff;
	    width: 500px;
	    height: 340px;
	    position: absolute;
	    top: 25%;
	    left: 35%;
	    padding: 20px;
			padding-right: 80px;
	}
	.login-title{
		font-size: 16px;
		font-weight: 400;
	}
	/*.login-wrap >>> .text-right{
		text-align: right;
	}*/
	.login-wrap >>> .text{
		font-size: 0.8em;
		font-weight: 60px;
		text-decoration-line: none;
	}
</style>
