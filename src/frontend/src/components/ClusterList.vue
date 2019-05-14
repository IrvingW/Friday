<template>
  <el-container>
      <el-container>
          <el-header>
              <!-- Header content -->
            <el-breadcrumb separator-class="el-icon-arrow-right">
              <el-breadcrumb-item :to="{ path: '/friday/home/cluster_list' }">集群列表</el-breadcrumb-item>
              <el-breadcrumb-item> </el-breadcrumb-item>
            </el-breadcrumb>
          </el-header>
          <el-main>
              <!-- Main content -->
              <el-card shadow="hover" :body-style="{ padding: '20px' }">
                <el-row class="button-row">
                  <el-col :span="2">
                    <el-button type="success" @click="addCluster" round icon="el-icon-folder-add" >添加集群</el-button>
                  </el-col>
                </el-row>
                <el-table :data="cluster_list" stripe>
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
                            disable-transitions>{{scope.row.status === 'error' ? '异常' : '健康'}}</el-tag>
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
                        <el-button size="normal" type="text" @click="showTaskList(scope.row.id)">任务列表</el-button>
                        <el-button size="normal" type="text" @click="showNodeView(scope.$index, scope.row)">节点视图</el-button>
                      </template>
                    </el-table-column>
                </el-table>
              </el-card>
          </el-main>
      </el-container>
  </el-container>
</template>
<script>
  export default {
	name:'cluster_list',
	  data(){
		return{
            columns: [
                {id: "name", label: "集群名称", "width": 150},
                {id: "create_time", label: "创建时间", "width": 150},
                {id: "description", label: "描述", "width": 330},
                {id: "node_cnt", label: "节点数量", "width": 120},
                {id: "submit_task", label: "历史任务", "width": 120},
                {id: "running_task", label: "正在运行", "width": 120},
            ],
            cluster_list: [
                {id: 1, name: "dclab", create_time: "2019-5-3", description: "分布式实验室集群", node_cnt: 2, submit_task: 1, running_task: 0, status: "健康"},
                {id: 1, name: "dclab", create_time: "2019-5-3", description: "分布式实验室集群", node_cnt: 2, submit_task: 1, running_task: 0, status: "健康"},
                {id: 1, name: "dclab", create_time: "2019-5-3", description: "分布式实验室集群", node_cnt: 2, submit_task: 1, running_task: 0, status: "健康"},
                {id: 1, name: "dclab", create_time: "2019-5-3", description: "分布式实验室集群", node_cnt: 2, submit_task: 1, running_task: 0, status: "健康"},
            ]
		}
    },
    methods: {
        showTaskList(cluster_id) {
            this.$router.push('/friday/home/task_list')
        },
        showNodeView(cluster_id) {
            this.$router.push('/friday/home/node_view')
        },
        addCluster() {
          this.$router.push('/friday/home/add_cluster')
        }
    }
  }
</script>
<style type="text/css" scoped>
.button-row {
  margin-top: 10px;
  margin-bottom: 20px;
  margin-left: 20px;
}
</style>