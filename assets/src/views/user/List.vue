<template>
  <page-header-wrapper :title="false">
    <a-card :body-style="{padding: '10px 10px'}" :bordered="false">
      <div class="table-operator">
        <a-button type="primary" icon="plus" @click="showModal">新建</a-button>
      </div>

      <add-form
        ref="addForm"
        :visible="visible"
        @cancel="handleCancel"
        @create="handleCreate"
      />

      <a-table
        :columns="columns"
        :data-source="items"
        rowKey="id"
        :pagination="pagination"
        @change="handlePageChange"
        :loading="loading"
      >
        <span slot="created" slot-scope="text, record">
          {{ record.created | formatTime }}
        </span>
        <span slot="updated" slot-scope="text, record">
          {{ record.updated | formatTime }}
        </span>
        <span slot="action" slot-scope="text, record">
          <template>
            <a-popconfirm
              title="确定删除该用户吗"
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
import { getUserList, deleteUser, addUser } from '@/api/user'
import AddForm from './AddForm.vue'

export default {
  name: 'TaskList',
  components: {
    AddForm
  },
  data () {
    return {
      // 查询参数
      queryParam: {},
      visible: false,
      // 表头
      columns: [
        {
          title: '用户ID',
          dataIndex: 'id',
          align: 'center',
          key: 'id'
        },
        {
          title: '用户名称',
          dataIndex: 'username',
          align: 'center',
          width: '200px'
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
    this.userList()
  },
  methods: {
    handleResetPwd (record) {
      const form = this.$refs.resetPwd.form
      form.validateFields((err, values) => {
        if (err) {
          return
        }
        addUser(values).then(() => {
          form.resetFields()
          this.resetPwdVisible = false
          this.userList()
        })
      })
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

    userList () {
      this.loading = true
      getUserList(this.queryParam).then(res => {
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
      this.userList()
    },

    showModal () {
      this.visible = true
    },

    handleCancel () {
      this.visible = false
    },

    handleCreate () {
      const form = this.$refs.addForm.form
      form.validateFields((err, values) => {
        if (err) {
          return
        }
        addUser(values).then((res) => {
          form.resetFields()
          this.visible = false
          this.userList()
        })
      })
    },

    handleDelete (record) {
      const params = {
        id: record.id
      }
      deleteUser(params).then((res) => {
        this.userList()
      })
    }
  }
}
</script>
