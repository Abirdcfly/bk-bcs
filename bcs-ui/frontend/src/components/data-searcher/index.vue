<template>
    <div class="biz-data-searcher">
        <template v-if="localScopeList.length">
            <template v-if="scopeDisabled">
                <bk-button class="bk-button trigger-btn disabled" style="max-width: 200px;">
                    <span class="btn-text tc">{{curScopeName}}</span>
                </bk-button>
            </template>
            <template v-else>
                <ClusterSelect style="width: 220px; float: left; position: relative; right: -1px; z-index: 10;"
                    v-model="localSearchScope"
                    searchable
                    :disabled="clusterFixed"
                    @change="handleSechScope"
                ></ClusterSelect>
            </template>
        </template>
        <div class="biz-search-input" style="width: 300px;">
            <bkbcs-input v-model="localKey"
                clearable
                :placeholder="placeholderRender"
                :right-icon="'bk-icon icon-search'"
                @enter="handleSearch"
                @clear="clearSearch">
            </bkbcs-input>
        </div>
        <div class="biz-refresh-wrapper" v-if="widthRefresh">
            <bk-popover class="refresh" :content="$t('刷新')" :delay="500" placement="top">
                <bk-button :class="['bk-button bk-default is-outline is-icon']" icon="icon-refresh" @click="handleRefresh">
                </bk-button>
            </bk-popover>
        </div>
    </div>
</template>

<script>
    //  -- 此组件已经和集群列表强相关 -- !!!注意!!!，有其他不是集群场景的组件，请单独实现
    import ClusterSelect from '@/components/cluster-selector/cluster-select.vue'

    export default {
        components: {
            ClusterSelect
        },
        props: {
            placeholder: {
                type: String,
                default: ''
            },
            searchPlaceholder: {
                type: String,
                default: ''
            },
            searchKey: {
                type: String,
                default: ''
            },
            searchScope: {
                type: String,
                default: ''
            },
            widthRefresh: {
                type: Boolean,
                default: true
            },
            scopeList: {
                type: Array,
                default () {
                    return []
                }
            },
            clusterFixed: {
                type: Boolean,
                default: false
            },
            scopeDisabled: {
                type: Boolean,
                default: false
            },
            searchable: {
                type: Boolean,
                default: true
            }
        },
        data () {
            return {
                isTriggerSearch: false,
                isRefresh: false,
                localKey: this.searchKey,
                localScopeList: [],
                placeholderRender: '',
                keyword: '',
                localSearchScope: '',
                curScopeId: ''
            }
        },
        computed: {
            curScopeName () {
                const curScope = this.scopeList.find(item => item.id === this.curScopeId)
                return curScope.name
            }
        },
        watch: {
            searchKey (val) {
                this.localKey = val
            },
            scopeList () {
                this.initLocalScopeList()
            },
            localKey (newVal, oldVal) {
                // 如果删除，为空时触发搜索
                if (oldVal && !newVal && !this.isRefresh) {
                    this.clearSearch()
                }
            },
            searchScope: {
                immediate: true,
                handler (val) {
                    if (val) {
                        this.localSearchScope = val
                        // this.handleSearch()
                    }
                }
            }
        },
        created () {
            this.initLocalScopeList()
            this.placeholderRender = this.placeholder || this.$t('输入关键字，按Enter搜索')
        },
        methods: {
            handleSechScope (id) {
                this.curScopeId = id
                this.handleSearch()
            },
            initLocalScopeList () {
                this.localScopeList = JSON.parse(JSON.stringify(this.scopeList))
                if (this.localScopeList.length) {
                    // 在初始化时，如果已经有值，选中
                    const clusterId = this.localSearchScope || sessionStorage['bcs-cluster']
                    if (clusterId) {
                        this.curScopeId = clusterId
                    } else {
                        this.curScopeId = this.localScopeList[0].id
                    }

                    sessionStorage['bcs-cluster'] = this.curScopeId
                    this.$emit('update:searchScope', this.curScopeId)
                }
            },
            handleSearch () {
                this.isTriggerSearch = true
                this.$emit('update:searchScope', this.curScopeId)
                this.$emit('update:searchKey', this.localKey)
                this.$emit('search')
                this.isRefresh = false
            },
            handleRefresh () {
                this.isRefresh = true
                this.$emit('refresh')
                this.$emit('update:searchKey', this.localKey)
            },
            clearSearch () {
                this.localKey = ''
                if (this.isTriggerSearch) {
                    this.handleSearch()
                    this.isTriggerSearch = false
                }
            }
        }
    }
</script>

<style scoped lang="postcss">
    @import '../../css/mixins/clearfix.css';
    @import '../../css/mixins/ellipsis.css';
    .biz-data-searcher {
        @mixin clearfix;

        .biz-search-input {
            .bk-form-input {
                border-radius: 0 2px 2px 0;
            }
        }

        .bk-dropdown-menu {
            .dropdown-item {
                > a {
                    width: 100%;
                    cursor: pointer;
                    display: inline-block;
                    vertical-align: middle;
                    @mixin ellipsis 240px;
                }
            }
            .bk-button {
                border-radius: 2px 0 0 2px;
                border-right: none;
            }
            float: left;
        }
    }
    .trigger-btn {
        &.disabled {
            margin-right: -10px;
            cursor: default;
            background: #fafafa;
        }
    }
    .btn-text {
        width: 140px;
        text-align: left;
        display: inline-block;
        vertical-align: middle;
        @mixin ellipsis 150px;
    }
</style>
