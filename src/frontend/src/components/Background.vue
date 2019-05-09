<template>
		<div  :style="background">
		</div>
</template>
<script>
	export default {
		name:'background',
		data(){
			return{
				background:{
					// backgroundImage:"url("+require("../assets/image/background.jpg")+")",
                    backgroundColor:'#f6f6f6',
					backgroundRepeat: "no-repeat",
					backgroundSize: "100% 100%",
					backgroundPosition: "center",
					width: "100vw",
					height: "100vh",
					position: "absolute",
					top: "0px",
					left: "0px",
					zIndex: "-1",
				}
			}
		}
	}
</script>
<style type="text/css" scoped>
</style>
