<template>
	<el-row type="flex" justify="space-between" class="header">
		<el-col :span="6"><div class="grid-content brand"> <router-link to="/">能源大数据采集管理系统</router-link></div></el-col>

		<el-col :span="6">
			<div class="grid-content lgout-options">
				<div>
					<el-link icon="el-icon-view">欢迎您，{{user_name}}   </el-link>
					<el-link icon="el-icon-back">退出登陆</el-link>
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
				user_name: ""
  		}
  	},
  	methods:{
			logout(){

        this.$router.push({path:'/login'})
			}
		},
    created(){
			var loginUser = ""
      var cookies = document.cookie.split(';');
      for(var i=0; i<cookies.length; i++)
      {
        var cookie = cookies[i].trim();
	      if (cookie.indexOf("user_name")==0)
	       loginUser = c.substring("user_name".length,cookie.length);
      }

      if(loginUser==""){
				// not login
				// this.$http.get("https://jsonplaceholder.typicode.com/todos/1")
				// 	.then(response => response.json())
				// 	.then(body => {
				// 		body.userId
				// })
        // this.$router.push({path:'/login'})
       }else{
				this.user_name = loginUser
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
