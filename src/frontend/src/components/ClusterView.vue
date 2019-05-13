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
                <el-button type="success" round icon="el-icon-s-order">显示集群列表</el-button>
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
          columns: ['日期', '集群数量', '机器数量', '执行任务数量'],
          rows: [
          ]
        },
        today: {
          columns: ['状态', '集群数量'],
          rows: [
            {'状态': "健康", '集群数量': 1},
            {'状态': "未定", '集群数量': 1},
            {'状态': "异常", '集群数量': 1}
            ]
        },
        tableData: [
          {name: "dclab", create_time: "2019-5-3", description: "分布式实验室集群", node_cnt: 2, submit_task: 1}
        ]
		  }
    },
    created: function() {
    }
  }
</script>
<style type="text/css" scoped>
</style>