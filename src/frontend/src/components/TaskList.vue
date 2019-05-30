<template>
  <el-container>
      <el-container>
          <el-header>
              <!-- Header content -->
            <el-breadcrumb separator-class="el-icon-arrow-right">
              <el-breadcrumb-item :to="{ path: '/friday/home/task_list' }">任务列表</el-breadcrumb-item>
              <el-breadcrumb-item> </el-breadcrumb-item>
            </el-breadcrumb>
          </el-header>
          <el-main>
              <!-- Main content -->
              <el-card shadow="hover" :body-style="{ padding: '20px' }">
                <el-row class="select-row">
                </el-row>
                <el-row>
                  <el-col :span="24">
                    <el-table :data="task_list" stripe>
                        <el-table-column v-for="col in columns"
                            :prop="col.id"
                            :key="col.id"
                            :label="col.label"
                            :width="col.width"
                            :sortable="col.id === 'create_time' || col.id === 'node_cnt' || col.id === 'submit_task' || col.id === 'running_task'">
                        </el-table-column>
                        <el-table-column label="状态" width="120">
                            <template slot-scope="scope">
                              <el-tag :type="scope.row.status === 'error' ? 'error' : 'success'"
                                disable-transitions>{{scope.row.status === 'error' ? '异常' : '完成'}}</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column width="300">
                          <template slot="header">
                            <el-input
                              v-model="searchInput"
                              size="normal"
                              placeholder="输入关键字搜索"/>
                          </template>
                          <template slot-scope="scope">
                            <el-button size="normal" type="text">停止</el-button>
                            <el-button size="normal" type="text" @click="showResult(scope.$index, scope.row)">查看结果</el-button>
                          </template>
                        </el-table-column>
                    </el-table>
                  </el-col>
                </el-row>
              </el-card>

              <el-dialog title="任务日志" :visible.sync="result_visible">
                <el-input
                  type="textarea" autosize
                  v-model="result_log">
                </el-input>
                  <div slot="footer" class="dialog-footer">
                    <el-button type="primary" @click="result_visible = false">确 定</el-button>
                  </div>
              </el-dialog>
          </el-main>
      </el-container>
  </el-container>
</template>
<script>
  export default {
	name:'task_list',
	  data(){
		return{
            columns: [
              {id: "name", label: "任务名称", "width": 140},
              {id: "creator", label: "创建者", "width": 140},
              {id: "cluster", label: "运行集群", "width": 140},
              {id: "create_time", label: "创建时间", "width": 140},
              {id: "description", label: "任务描述", "width": 300},
              //{id: "status", label: "状态", "width": 140},
            ],
            task_list: [
            ],
            searchTaskUrl: '/api/task/search?key=',
            result_visible: false,
            result_log: ""
		}
    },
    methods: {
        showResult(index, row) {
          console.log(row.output)
          this.result_log = row.output
          this.result_visible = true
        }
    },
    created: function () {
      this.$http.get(this.searchTaskUrl, {withCredentials: true}).then((response) => {
			  if (response.status != 200){
				  this.$message.error({
					  message: '请求发送失败，请联系管理员'
				  })
			  }
			  var body = response.body
			  if (body.Rtn == 0){
          this.task_list = body.Tasks
          this.task_list.forEach(element => {
            element.create_time = element.create_time.split("T")[0]
            element.creator = element.creator.user_name
            element.cluster = element.cluster.name
          });
			  }else {
				  this.$message.error({
                message: body.Msg
              });
            }
      })
    }
  }
</script>
<style type="text/css" scoped>
.select {
  margin-top: 10px;
}
.select-row {
  margin-bottom: 30px;
}
</style>