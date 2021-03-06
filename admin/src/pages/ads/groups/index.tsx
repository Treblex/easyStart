import List from '@/components/List';
import PageMain from '@/components/PageMain';
import { useDataList } from '@/hooks/useDataList';
import { adGroups } from '@/server/api/ad';
import { DeleteOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, Drawer, Modal, Space, Tooltip, Typography } from 'antd';
import React, { useState } from 'react';
import { AddAdGroup } from './components/add';
import { Detail } from './components/detail';
let { confirm } = Modal;
const { Paragraph } = Typography;
export default function Groups() {
  let { data, load, loading } = useDataList(
    page => adGroups.list({ page }),
    true,
  );
  // 编辑分组
  let [visible, setVisible] = useState(false);

  let [chooseId, setChooseId] = useState(0);

  // 查看详情
  let [detailVisible, setDetailVisible] = useState(false);

  let values: { name: string; value: any }[] = [];
  let [selectValues, setSelectValues] = useState(values);

  const deleteGroup = (id: number) => {
    confirm({
      title: '确定要删除广告位？无法撤回',
      icon: <DeleteOutlined color="#d33"></DeleteOutlined>,
      okText: '确定删除',
      okType: 'danger',
      onOk() {
        adGroups.del(id).then(() => {
          load();
          setSelectValues([]);
        });
      },
    });
  };

  const columns = [
    { title: 'ID', key: 'id', dataIndex: 'id' },
    { title: '广告位标题', key: 'name', dataIndex: 'name' },

    {
      title: '介绍',
      key: 'desc',
      dataIndex: 'desc',
      width: 400,
      rows: 2,
      render: (desc: React.ReactNode) => (
        <Paragraph ellipsis={{ rows: 2, expandable: false }}>{desc}</Paragraph>
      ),
    },
    {
      title: '统计',
      key: 'count',
      dataIndex: 'count',
      render: (count: number, record: any) => {
        return (
          <Space>
            <span>{`(${count}/${record.max_count})`}</span>
            <a
              onClick={() => {
                setDetailVisible(true);
                setChooseId(record.id);
              }}
            >
              添加广告
            </a>
          </Space>
        );
      },
    },
    {
      title: '更新时间',
      key: 'updated_at',
      dataIndex: 'updated_at',
    },
    {
      title: '创建时间',
      key: 'created_at',
      dataIndex: 'created_at',
    },
    {
      title: '操作',
      key: 'action',
      dataIndex: 'id',
      render: (id: number, data: any) => {
        return (
          <Space style={{ width: 120 }}>
            <a
              onClick={() => {
                setSelectValues(
                  Object.keys(data).map(x => {
                    return { name: x, value: data[x] };
                  }),
                );
                setVisible(true);
              }}
            >
              编辑
            </a>
            <span> / </span>
            {(() => {
              if (data.count > 0 || data.tag_count > 0) {
                return (
                  <Tooltip title="该广告位下有子内容，不可删除，请先清理广告">
                    <span>不可删除</span>
                  </Tooltip>
                );
              } else {
                return <a onClick={() => deleteGroup(data.id)}>删除</a>;
              }
            })()}
          </Space>
        );
      },
    },
  ];

  const leftActions = () => [
    <Button
      key="add-group"
      type="primary"
      onClick={() => {
        setSelectValues([
          { name: 'name', value: '' },
          { name: 'desc', value: '' },
          { name: 'max_count', value: 3 },
        ]);
        setVisible(true);
      }}
    >
      <PlusOutlined />
      添加广告位
    </Button>,
  ];
  return (
    <PageMain
      title="广告位管理"
      subTitle="全局通用的banner 海报管理,最大数量用于前端判断轮播图或者固定位海报，这里后台暂时没有限制最大数量，由前端处理"
    >
      <Drawer
        width={800}
        visible={visible}
        onClose={() => {
          setVisible(false);
          load();
        }}
      >
        <AddAdGroup
          values={selectValues}
          callback={() => {
            setVisible(false);
            load();
          }}
        ></AddAdGroup>
      </Drawer>

      <Drawer
        width={1000}
        visible={detailVisible}
        onClose={() => {
          setDetailVisible(false);
          load();
        }}
      >
        <Detail id={chooseId}></Detail>
      </Drawer>

      <List
        leftActions={leftActions()}
        onRefresh={load}
        loading={loading}
        table={{
          columns,
          dataSource: data.list,
          bordered: true,
          rowKey: 'id',
          loading,
          pagination: {
            position: ['bottomLeft'],
            total: data.total,
          },
        }}
      ></List>
    </PageMain>
  );
}
