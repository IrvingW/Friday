<template>
	<el-row type="flex" justify="space-between" class="header">
		<el-col :span="6"><div class="grid-content brand"> <router-link to="/">能源大数据处理任务调度管理系统</router-link></div></el-col>

		<el-col :span="6">
			<div class="grid-content lgout-options">
				<div>
					<el-link icon="el-icon-user-solid" @click="showUserView">欢迎您，{{user_name}}         </el-link>
					<el-link icon="el-icon-s-promotion" @click="logout">退出登陆</el-link>
				</div>

			</div>
		</el-col>

	</el-row>
</template>
<script>
  export default {
  	name:"navmenu",
  	data(){
  		return {
				user_name: "",
				logout_url: "/api/user/logout",
				if_login_url: "/api/user/if_login"
  		}
  	},
  	methods:{
			logout(){
				this.$http.get(this.logout_url, {withCredentials: true}).then(
					(response) => {
					  if (response.status != 200){
						  this.$message.error({
							  message: '请求发送失败，请联系管理员'
						  })
					  }else{
						  var body = response.body
						  if (body.Rtn == 0){
							  this.$message({
									 message: '退出登陆成功',
                   type: 'success'
							  });
							  // set cookie
							  this.$cookies.remove("user_name")
							  // redirect to index page
							  this.$router.push({path: "/friday/login"})
						  }else{
							  this.$message.error({
                   message: body.Msg
							   })
						  }
						}
					}
				)
			},
			showUserView(){
        this.$router.push('/friday/home/user_view')
			}
    },
		created: function() {
       if(!this.$cookies.isKey("user_name")){
				 console.log(this.if_login_url)
				 this.$http.get(this.if_login_url, {withCredentials: true}).then(
					 (response) => {
				 console.log("haha")
					   if (response.status != 200){
						   this.$message.error({
							   message: '登陆信息获取失败，请联系管理员'
						   })
					   }else{
						   var body = response.body
						   if (body.Rtn == 0){
								 if (body.Username == ""){
							     // remove cookie
							     this.$cookies.remove("user_name")
							     // redirect to index page
									this.$router.push({path: "/friday/login"})
								 }else {
									 //logged in but cookie is expired
							     // reset cookie
							     this.$cookies.set("user_name", body.Username, 600)
								 }
						   }else{
						   }
						 }
					}
				)
			} else{
				 this.user_name = this.$cookies.get("user_name")
			 }
     }

  }
</script>

<style scoped="scoped">
	.header{
	  	height: 60px;
	  	line-height: 60px;
	  	background-color: rgb(6,108,154);
	 }
	.brand,.lgout-options {
	  	color: white;
	 }
	.brand {
	  	padding-left: 20px;
	 }
	.lgout-options{
	  	text-align: right;
	  	padding-right: 40px;

	}
  	.el-row {
      margin-bottom: 20px;
    }

	.grid-content {
	  height: 60px;
	  line-height: 60px;

	}
	.grid-content ,.grid-content a{
		color: white;
		height: 60px;
		display: inline-block;
	}
	.el-button:first-child,a:first-child{
		margin-right: 8px;
	}
	a:hover,a{
		text-decoration-line: none;
	}

</style>
