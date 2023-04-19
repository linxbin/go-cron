<template>
  <page-header-wrapper :title="false">
    <a-card :body-style="{padding: '10px 10px'}" :bordered="false">
      <div class="table-operator">
        <a-button type="primary" icon="plus" @click="goAdd">新建</a-button>
      </div>
      <div class="table-page-search-wrapper">
        <a-form layout="inline">
          <a-row :gutter="48">
            <a-col :md="4" :sm="12">
              <a-form-item label="任务名称">
                <a-input v-model="queryParam.name" placeholder=""/>
              </a-form-item>
            </a-col>
            <a-col :md="4" :sm="12">
              <a-form-item label="使用状态">
                <a-select v-model="queryParam.status" placeholder="请选择" default-value="0">
                  <a-select-option value="0">全部</a-select-option>
                  <a-select-option value="10">运行中</a-select-option>
                  <a-select-option value="20">停止</a-select-option>
                </a-select>
              </a-form-item>
            </a-col>
            <a-col :md="2" :sm="6">
              <span class="table-page-search-submitButtons">
                <a-button type="primary" @click="taskList">查询</a-button>
                <a-button style="margin-left: 8px" @click="() => queryParam = {}">重置</a-button>
              </span>
            </a-col>
          </a-row>
        </a-form>
      </div>

      <a-table
        :columns="columns"
        :data-source="items"
        rowKey="id"
        :pagination="pagination"
        @change="handlePageChange"
        :loading="loading"
      >
        <span slot="status" slot-scope="text, record">
          <a-switch checked-children="开" un-checked-children="关" @change="handleChangeStatus(record)" :default-checked="record.is_enable" />
        </span>
        <span slot="created" slot-scope="text, record">
          {{ record.created | formatTime }}
        </span>
        <span slot="updated" slot-scope="text, record">
          {{ record.updated | formatTime }}
        </span>
        <span slot="action" slot-scope="text, record">
          <template>
            <a @click="goTaskLog(record.id)">查看日志</a>
            <a-divider type="vertical" />
            <a @click="handleEdit(record)">编辑</a>
            <a-divider type="vertical" />
            <a-popconfirm
              title="确定删除该任务吗"
              ok-text="确定"
              cancel-text="取消"
              @confirm="handleDelete(record)"
            >
              <a href="#">删除</a>
            </a-popconfirm>

          </template>
        </span>
      </a-table>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import moment from 'moment'
import { getTaskList, enableTask, disableTask, deleteTask } from '@/api/task'

export default {
  name: 'TaskList',
  components: {
  },
  data () {
    return {
      // 查询参数
      queryParam: {},
      // 表头
      columns: [
        {
          title: '任务ID',
          dataIndex: 'id',
          align: 'center',
          key: 'id'
        },
        {
          title: '任务名称',
          dataIndex: 'name',
          align: 'center',
          width: '200px'
        },
        {
          title: 'cron表达式',
          dataIndex: 'spec',
          align: 'center',
          width: '200px'
        },
        {
          title: '状态',
          dataIndex: 'status',
          align: 'center',
          scopedSlots: { customRender: 'status' }
        },
        {
          title: '创建时间',
          dataIndex: 'created',
          align: 'center',
          scopedSlots: { customRender: 'created' }
        },
        {
          title: '更新时间',
          dataIndex: 'updated',
          align: 'center',
          scopedSlots: { customRender: 'updated' }
        },
        {
          title: '操作',
          dataIndex: 'action',
          align: 'center',
          width: '250px',
          scopedSlots: { customRender: 'action' }
        }
      ],
      items: [],
      selectedRowKeys: [],
      selectedRows: [],
      loading: false,
      pagination: {
        current: 1,
        total: 0,
        pageSize: 0
      }
    }
  },
  created () {
    this.taskList()
  },
  methods: {
    handleEdit (record) {
      this.$router.push({ name: 'TaskEdit', params: { id: record.id } })
    },

    onSelectChange (selectedRowKeys, selectedRows) {
      this.selectedRowKeys = selectedRowKeys
      this.selectedRows = selectedRows
    },

    resetSearchForm () {
      this.queryParam = {
        date: moment(new Date())
      }
    },

    taskList () {
      this.loading = true
      getTaskList(this.queryParam).then(res => {
        this.items = res.result.items
        this.pagination.current = res.result.pager.page
        this.pagination.total = res.result.pager.total_rows
        this.pagination.pageSize = res.result.pager.page_size
        this.loading = false
      })
    },

    handlePageChange (pagination) {
      this.pagination.current = pagination.current
      this.queryParam.page = pagination.current
      this.taskList()
    },

    handleChangeStatus (record) {
      const params = {
        id: record.id
      }
      if (record.is_enable) {
        this.handleDisableTask(params)
      } else {
        this.handleEnableTask(params)
      }
    },
    handleDisableTask (params) {
      disableTask(params).then((res) => {
        this.taskList()
      })
    },
    handleEnableTask (params) {
      enableTask(params).then((res) => {
        this.taskList()
      })
    },
    handleDelete (record) {
      const params = {
        id: record.id
      }
      deleteTask(params).then((res) => {
        this.taskList()
      })
    },

    goTaskLog (id) {
      this.$router.push({ path: `/task-log/list/${id}` })
    },

    goAdd () {
      this.$router.push({ name: 'TaskAdd' })
    }
  }
}
</script>
