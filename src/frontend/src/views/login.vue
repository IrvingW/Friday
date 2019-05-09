<template>
	<el-container>
    <el-main>
	  <div class="login-wrap">
		  <Background></Background>
		  <div id="login">
			  <h2 class="login-title">用户登录</h2>
			  <el-form label-position="top" label-width="100px" :model="login" :rules="rules" ref="login">
			    <el-form-item label="用户名" prop="username">
			      <el-input v-model="login.username" clearable></el-input>
			    </el-form-item>
			    <el-form-item label="密码" prop="password">
			      <el-input v-model="login.password" type="password"></el-input>
			    </el-form-item>
					<el-divider><i class="el-icon-lock"></i></el-divider>
			    <el-row type="flex" justify="space-between">
			  	  <el-col :span="12" class="text-right">
			  		  <router-link to="/register" class="el-link" type="primary">没有账号，前往注册</router-link>
			  	  </el-col>
			    </el-row>
			    <el-form-item>
			        <el-button size="medium" @click="logIn">登录</el-button>
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
	      return {
	        login: {
	          username: '',
	          password: '',
	        },
	        rules:{
	        	username:[
	        		{required:true,message:'请输入用户名',trigger:'blur'},
	        		{min:1,message:'用户名不能为空',trigger:'blur'}
	        	],
	        	password:[
	        		{required:true,message:'请输入密码',trigger:'blur'},
	        		{min:1,message:'密码不能为空',trigger:'blur'}
	        	]
	        }
	      }
		},
		methods: {
			logIn(){
				this.login.username
			},

	  }
	}
</script>



<style scoped>
	.login-wrap >>> .el-form-item__label {
		margin-bottom: -7px;
		padding-bottom: 0px;
	}
	.login-wrap >>> .el-form-item:first-child{
		margin-bottom: 10px;
	}
	.login-wrap >>> .el-form-item:not(:first-child){
		margin-bottom: 15px;
	}
	.login-wrap >>> .el-row {
		margin-bottom: 10px;
	}
	.login-wrap >>> .el-button{
		width: 100%;
	}
	.login-wrap >>> label {
		font-size: 13px;
		font-weight: 200;
	}
	#login {
	    background-color: #ffffff;
	    width: 300px;
	    height: 340px;
	    position: absolute;
	    top: 25%;
	    left: 40%;
	    padding: 20px;
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
