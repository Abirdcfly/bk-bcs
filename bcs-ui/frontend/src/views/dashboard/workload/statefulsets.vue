<template>
    <BaseLayout title="StatefulSets" kind="StatefulSet" category="statefulsets" type="workloads">
        <template #default="{ curPageData, pageConf, handlePageChange, handlePageSizeChange, handleGetExtData, gotoDetail, handleSortChange,handleUpdateResource,handleDeleteResource }">
            <bk-table
                :data="curPageData"
                :pagination="pageConf"
                @page-change="handlePageChange"
                @page-limit-change="handlePageSizeChange"
                @sort-change="handleSortChange">
                <bk-table-column :label="$t('名称')" prop="metadata.name" sortable>
                    <template #default="{ row }">
                        <bk-button class="bcs-button-ellipsis" text @click="gotoDetail(row)">{{ row.metadata.name }}</bk-button>
                    </template>
                </bk-table-column>
                <bk-table-column :label="$t('命名空间')" prop="metadata.namespace" sortable></bk-table-column>
                <bk-table-column :label="$t('镜像')" prop="extraFE.images" width="450" :show-overflow-tooltip="false">
                    <template slot-scope="{ row }">
                        <span v-bk-tooltips.top="(handleGetExtData(row.metadata.uid, 'images') || []).join('<br />')">
                            {{ (handleGetExtData(row.metadata.uid, 'images') || []).join(', ') }}
                        </span>
                    </template>
                </bk-table-column>
                <bk-table-column label="Ready" width="110" :resizable="false">
                    <template slot-scope="{ row }">{{row.status.readyReplicas || 0}} / {{row.spec.replicas || 0}}</template>
                </bk-table-column>
                <bk-table-column label="Up-to-date" width="110" :resizable="false">
                    <template slot-scope="{ row }">{{row.status.updatedReplicas || 0}}</template>
                </bk-table-column>
                <bk-table-column label="Age" prop="extraFE.age" :resizable="false">
                    <template #default="{ row }">
                        <span>{{handleGetExtData(row.metadata.uid, 'age')}}</span>
                    </template>
                </bk-table-column>
                <bk-table-column :label="$t('编辑模式')" width="100">
                    <template slot-scope="{ row }">
                        <span>
                            {{handleGetExtData(row.metadata.uid, 'editMode') === 'form'
                                ? $t('表单') : 'YAML'}}
                        </span>
                    </template>
                </bk-table-column>
                <bk-table-column :label="$t('操作')" :resizable="false" width="150">
                    <template #default="{ row }">
                        <bk-button text
                            @click="handleUpdateResource(row)">{{ $t('更新') }}</bk-button>
                        <bk-button class="ml10" text
                            @click="handleDeleteResource(row)">{{ $t('删除') }}</bk-button>
                    </template>
                </bk-table-column>
            </bk-table>
        </template>
    </BaseLayout>
</template>
<script>
    import { defineComponent } from '@vue/composition-api'
    import BaseLayout from '@/views/dashboard/common/base-layout'

    export default defineComponent({
        components: { BaseLayout }
    })
</script>
