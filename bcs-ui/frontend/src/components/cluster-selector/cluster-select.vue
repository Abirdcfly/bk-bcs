<template>
    <bcs-select class="cluster-select"
        v-model="localValue"
        :clearable="false"
        :searchable="searchable"
        :disabled="disabled"
        @change="handleClusterChange">
        <bcs-option v-for="item in clusterList"
            :key="item.clusterID"
            :id="item.clusterID"
            :name="item.clusterName"
        ></bcs-option>
    </bcs-select>
</template>
<script>
    import { computed, defineComponent, ref, watch, toRefs } from '@vue/composition-api'
    import { BCS_CLUSTER } from '@/common/constant'
    import useDefaultClusterId from '@/views/node/use-default-clusterId'

    export default defineComponent({
        name: 'ClusterSelect',
        props: {
            value: {
                type: String,
                default: ''
            },
            searchable: {
                type: Boolean,
                default: false
            },
            disabled: {
                type: Boolean,
                default: false
            }
        },
        model: {
            prop: 'value',
            event: 'change'
        },
        setup (props, ctx) {
            const { $store } = ctx.root
            const { value } = toRefs(props)
            
            watch(value, (v) => {
                localValue.value = v
            })

            const clusterList = computed(() => {
                return $store.state.cluster.clusterList || []
            })
            // 集群ID默认规则：props > 单集群ID > sessionStorage > 取列表第一个
            const { defaultClusterId } = useDefaultClusterId()
            const localValue = ref(props.value || defaultClusterId.value)

            const handleClusterChange = (clusterId) => {
                localValue.value = clusterId
                sessionStorage.setItem(BCS_CLUSTER, clusterId)
                ctx.emit('change', clusterId)
            }
            return {
                localValue,
                clusterList,
                handleClusterChange
            }
        }
    })
</script>
<style lang="postcss" scoped>
.cluster-select {
    min-width: 200px;
    background-color: #fff;
}
</style>
