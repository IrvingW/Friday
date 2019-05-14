<template>
<el-container>
  <el-container>
    <el-header>
      <!-- Header content -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/friday/home/cluster_view' }">集群概览</el-breadcrumb-item>
        <el-breadcrumb-item> </el-breadcrumb-item>
      </el-breadcrumb>
    </el-header>
    <el-main>
      <!-- Main content -->
      <el-row>
        <el-col :span="23">
          <el-card shadow="hover" :body-style="{ padding: '20px' }">
            <!-- charts row -->
            <el-row :gutter="30">
              <el-col :span="8">
                <ve-pie :data="today"></ve-pie>
              </el-col>
              <el-col :span="15">
                <ve-line :data="history"></ve-line>
              </el-col>
            </el-row>
            <!-- list row -->
            <el-row :gutter="30">
              <el-col :span="18">
                <el-tabs>
                  <el-tab-pane label="执行任务Top3集群">
                      <el-table
                        :data="tableData"
                        stripe
                        style="width: 100%">
                        <el-table-column
                          prop="name"
                          label="集群名称"
                          width="180">
                        </el-table-column>
                        <el-table-column
                          prop="create_time"
                          label="创建时间"
                          width="180">
                        </el-table-column>
                        <el-table-column
                          prop="description"
                          label="描述"
                          width="180">
                        </el-table-column>
                        <el-table-column
                          prop="node_cnt"
                          label="节点数量"
                          width="180">
                        </el-table-column>
                        <el-table-column
                          prop="submit_task"
                          label="执行任务数量">
                        </el-table-column>
                      </el-table>
                  </el-tab-pane>
                </el-tabs>
              </el-col>
              <el-col :span="2" :offset="1">
                <el-row class="button_row">
                  <el-col>
                    <el-button type="primary" round icon="el-icon-s-order" @click="showClusterList">集群列表</el-button>
                  </el-col>
                </el-row>
                <el-row class="button_row">
                  <el-col>
                    <el-button type="success" round icon="el-icon-folder-add" @click="addCluster">添加集群</el-button>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</el-container>

</template>
<script>
  export default {
	name:'cluster_view',
	  data(){
		  return{
        history: {
          columns: ['date', '集群数量', '机器数量', '执行任务数'],
          rows: [
            {'date': "5-5", '执行任务数': '3', '集群数量': 1, '机器数量':2},
            {'date': "5-6", '执行任务数': '2', '集群数量': 1, '机器数量':2},
            {'date': "5-7", '执行任务数': '4', '集群数量': 1, '机器数量':2},
            {'date': "5-8", '执行任务数': '5', '集群数量': 1, '机器数量':2},
            {'date': "5-9", '执行任务数': '2', '集群数量': 1, '机器数量':2},
            {'date': "5-10", '执行任务数': '4', '集群数量': 1, '机器数量':2},
          ]
        },
        today: {
          columns: ['status', 'cnt'],
          rows: [
            {'status': "健康集群", 'cnt': 1},
            {'status': "未定", 'cnt': 1},
            {'status': "异常集群", 'cnt': 1}
            ]
        },
        tableData: [
          {name: "dclab", create_time: "2019-5-3", description: "分布式实验室集群", node_cnt: 2, submit_task: 1},
          {name: "dclab", create_time: "2019-5-3", description: "分布式实验室集群", node_cnt: 2, submit_task: 1},
          {name: "dclab", create_time: "2019-5-3", description: "分布式实验室集群", node_cnt: 2, submit_task: 1}
        ]
		  }
    },
    created: function() {
    },
    methods: {
      showClusterList(){
        this.$router.push('/friday/home/cluster_list')
      },
      addCluster(){
        this.$router.push('/friday/home/add_cluster')
      }
    }
  }
</script>
<style type="text/css" scoped>
.button_row {
  margin-top: 30px;
}
</style>