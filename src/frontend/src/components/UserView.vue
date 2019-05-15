<template>
<el-container>
  <el-container>
    <el-header>
      <!-- Header content -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
          <el-breadcrumb-item :to="{ path: '/friday/home/user_view' }">用户中心</el-breadcrumb-item>
          <el-breadcrumb-item> </el-breadcrumb-item>
        </el-breadcrumb>
    </el-header>
    <el-main>
      <!-- Main content -->
        <el-card shadow="hover" :body-style="{ padding: '20px' }">
          <el-tabs active-name="cluster">
            <el-tab-pane label="授权集群" name="cluster">
              <el-row>
                <el-col :span="4">
                  <p>用户名：{{user_name}}</p>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="1"> <p>角色：</p> </el-col>
                <el-col :span="2">
                  <el-tooltip class="item" effect="light" :content="tip" placement="right">
                    <el-button type="primary" plain>{{role === 'admin' ? '管理员' : '普通用户'}}</el-button>
                  </el-tooltip>
                </el-col>
                <el-col :span="1">
                  <el-button class="item" type="success" @click="createUser" icon="el-icon-s-custom">创建新用户</el-button>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="4">
                  <p>已授权集群列表：</p>
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
            </el-tab-pane>
            <el-tab-pane label="任务管理" name="task">
              <el-row>
                <el-col :span="4">
                  <p>用户名：{{user_name}}</p>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="1"> <p>角色：</p> </el-col>
                <el-col :span="2">
                  <el-tooltip class="item" effect="light" :content="tip" placement="right">
                    <el-button type="primary" plain>{{role === 'admin' ? '管理员' : '普通用户'}}</el-button>
                  </el-tooltip>
                </el-col>
                <el-col :span="1">
                  <el-button class="item" type="success" @click="createUser" icon="el-icon-s-custom">创建新用户</el-button>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="4">
                  <p>历史执行任务：</p>
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
            </el-tab-pane>
          </el-tabs>
        </el-card>
        <el-dialog title="选择授权给该用户使用集群" :visible.sync="dialogFormVisible">
          <el-form :model="form">
            <el-form-item label="授权集群" :label-width="formLabelWidth">
                <el-select v-model="form.clusters" multiple placeholder="请选择要授权的集群" :clearable="true" size="medium">
                  <el-option
                    v-for="cluster in clusters"
                    :key="cluster.id"
                    :label="cluster.name"
                    :value="cluster.id">
                  </el-option>
                </el-select>
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button @click="resetForm()">取 消</el-button>
            <el-button type="primary" @click="submitForm()">确 定</el-button>
          </div>
        </el-dialog>
    </el-main>
  </el-container>
</el-container>


</template>
<script>
  export default {
	name:'user_view',
	  data(){
		  return{
        form: {
          clusters: []
        },
        dialogFormVisible: false,
        clusters: [
          {id: 0, name: "dclab"},
          {id: 1, name: "dblab"},
          {id: 2, name: "dalab"}
        ],
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
        ],
        user_name: '',
        role: '',
        tip: '',
      }
    },
    created: function() {
      this.user_name = this.$cookies.get("user_name")
      this.role = 'admin'
      if (this.role === 'admin') {
        this.tip = '管理员可以给新用户授权，并拥有所有集群的权限'
      }else {
        this.tip = '普通用户只能在已授权的集群上创建任务'
      }
    },
    methods: {
      createUser() {
        if (this.role != 'admin'){
          this.$alert('只用管理员用户可以创建新用户', '无法创建新用户', {
            confirmButtonText: '确定',
            callback: action => {
              this.$message({
                type: 'info',
                message: `action: ${ action }`
              });
            }
          });
        }
        this.dialogFormVisible = true
      },
      resetForm() {
        this.dialogFormVisible = false;
      },
      submitForm() {
        this.dialogFormVisible = false;
      }
    }
  }
</script>
<style type="text/css" scoped>
.item {
  margin-top: 10px;
}
</style>