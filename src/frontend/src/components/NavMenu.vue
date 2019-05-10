<template>
	<el-row type="flex" justify="space-between" class="header">
		<el-col :span="6"><div class="grid-content brand"> <router-link to="/">能源大数据采集管理系统</router-link></div></el-col>

		<el-col :span="6">
			<div class="grid-content lgout-options">
				<div>
					<el-link icon="el-icon-view">欢迎您，{{user_name}}   </el-link>
					<el-link icon="el-icon-back" @click="logout">退出登陆</el-link>
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
				logout_url: "http://localhost:8080/api/user/logout"
  		}
  	},
  	methods:{
			logout(){
				this.$http.get(this.logout_url).then(
					(response) => {
					  if (response.status != 200){
						  this.$notify.error({
							  title: '请求失败',
							  message: '请求发送失败，请联系管理员'
						  })
					  }else{
						  var body = response.body
						  if (body.Rtn == 0){
							  this.$notify({
									 title: '成功',
									 message: '退出登陆成功',
                   type: 'success'
							  });
							  // set cookie
							  this.$cookies.remove("user_name")
							  // redirect to index page
							  this.$router.push({path: "/login"})
						  }else{
							  this.$notify.error({
								  title: '失败',
                   message: body.Msg
							   })
						  }
						}
					}
				)
			}
		},
      created(){
        if(!this.$cookies.isKey("user_name")){
				  // not login
          this.$router.push({path:'/login'})
         }else{
				  this.user_name = this.$cookies.get("user_name")
			  }
      },

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
	  	padding-right: 20px;

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
