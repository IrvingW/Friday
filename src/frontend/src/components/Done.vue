<template>
<el-container>
	<el-container>
		<el-header height="50px">
			<!-- Header content -->
		</el-header>
		<el-main>
			<!-- Main content -->
			<el-row>
				<el-col :span="6" :offset="8">
					<el-row>
					  <img src="../assets/icon-ok.png">
					</el-row>
					<el-row>
					  <p class="message">{{message}}</p>
					</el-row>
					<el-row>
					  <el-button type="primary" @click="showCLusters" class="right">完成</el-button>
					</el-row>
				</el-col>
			</el-row>
		</el-main>
	</el-container>
</el-container>

</template>
<script>
  export default {
	name:'done',
	props:{
		message:{
			type: String
		}
	},
	data(){
		return{
		}
	},
	methods: {
		showClusters(){
			this.$router.push('/friday/home/cluster_list')
		}
	}
  }
</script>
<style type="text/css" scoped>
.right {
	margin-left: 85px
}
.message {
	font-family: "Helvetica Neue";
	font-size: 20px;
	margin-left: 60px;
}
</style>