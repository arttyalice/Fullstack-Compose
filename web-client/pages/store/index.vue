<template>
  <div class="container">
    <c-box
        d="flex"
        justify-content="flex-end"
        p="2"
    >
        <c-button @click="createStoreModalOpen">Add Store</c-button>
    </c-box>
    <c-box
        p="5"
    >
        <c-simple-grid :columns="3" :spacing="2">
            <c-box
                p="3"
                border-width="1px"
                justify-content="center"
                v-for="(store, index) in list"
                :key="store.id"
            >
                <c-heading text-align="center" size="xl">
                    {{ store.name }}
                </c-heading>
                <c-flex justify="center" direction="column" align="center">
                    Description: {{ store.description }}<br />
                    Phone: {{ store.phone_number }}<br />
                    Address: {{ store.address }}
                </c-flex>
                <c-flex justify="center" direction="column">
                    <c-button variant-color="blue" as="a" :href="`/store/${store.id}`">
                        Info
                    </c-button>
                    <c-button variant-color="green" @click="editStoreModalOpen(store)">
                        Edit
                    </c-button>
                    <c-button @click="deleteStore(store.id, index)">
                        Delete
                    </c-button>
                </c-flex>
            </c-box>
        </c-simple-grid>
    </c-box>

    <!-- Create -->
    <c-modal
      :is-open="isCreateStore"
      :on-close="createStoreModalClose"
    >
        <c-modal-content>
            <c-modal-header>Create Store</c-modal-header>
            <c-modal-close-button />
            <c-modal-body>
                <c-form-control mt="4" is-required>
                    <c-form-label>Name</c-form-label>
                    <c-input v-model="newStore['name']" placeholder="Store Name" />
                </c-form-control>

                <c-form-control mt="4" is-required>
                    <c-form-label>Phone</c-form-label>
                    <c-input v-model="newStore['phone_number']" placeholder="Store Phone number" />
                </c-form-control>

                <c-form-control mt="4" is-required>
                    <c-form-label>Address</c-form-label>
                    <c-input v-model="newStore['address']" placeholder="Store Address" />
                </c-form-control>

                <c-form-control mt="4" is-required>
                    <c-form-label>Description</c-form-label>
                    <c-textarea v-model="newStore['description']" placeholder="Store Description" />
                </c-form-control>
            </c-modal-body>
            <c-modal-footer>
                <c-button @click="createStoreModalClose" mr="3">
                    Cancel
                </c-button>
                <c-button @click="createStore" variant-color="blue">Save</c-button>
            </c-modal-footer>
        </c-modal-content>
        <c-modal-overlay />
    </c-modal>

    <!-- Edit -->
    <c-modal
      :is-open="isEditStore"
      :on-close="editStoreModalClose"
    >
      <c-modal-content>
            <c-modal-header>Edit Product</c-modal-header>
            <c-modal-close-button />
                <c-modal-body>
                    <c-form-control mt="4" is-required>
                        <c-form-label>Name</c-form-label>
                        <c-input v-model="editStore['name']" placeholder="Store Name" />
                    </c-form-control>

                    <c-form-control mt="4" is-required>
                        <c-form-label>Phone</c-form-label>
                        <c-input v-model="editStore['phone_number']" placeholder="Store Phone number" />
                    </c-form-control>

                    <c-form-control mt="4" is-required>
                        <c-form-label>Address</c-form-label>
                        <c-input v-model="editStore['address']" placeholder="Store Address" />
                    </c-form-control>

                    <c-form-control mt="4" is-required>
                        <c-form-label>Description</c-form-label>
                        <c-textarea :value="editStore['description']" v-model="editStore['description']" placeholder="Store Description" />
                    </c-form-control>
                </c-modal-body>
            <c-modal-footer>
                <c-button @click="editStoreModalClose" variant-color="blue" mr="3">
                    Cancel
                </c-button>
                <c-button @click="updateStore">Save</c-button>
            </c-modal-footer>
        </c-modal-content>
        <c-modal-overlay />
    </c-modal>
  </div>
</template>

<script lang="js">

export default {
	name: 'App',
	data () {
		return {
            list: [],
            // create store data
            isCreateStore: false,
            newStore: {},
            // edit store data
            isEditStore: false,
            editStore: {},
		}
    },
    created () {
        this.fetchStores()
    },
    methods: {
        fetchStores () {
            this.$axios.get('/api/store/list')
            .then(res => {
                this.list = res.data.items
            })
            .catch(err => {
                console.log(err)
            })
        },
        async createStore () {
            if (
                !this.newStore.name ||
                !this.newStore.description ||
                !this.newStore.phone_number ||
                !this.newStore.address
            ) {
                this.$toast({
                    title: 'Please complete the form.',
                    description: "Please fill all required fields in the form.",
                    status: 'error',
                    duration: 10000
                })
            }
            await this.$axios.post('/api/store/create', this.newStore)
            this.fetchStores()
            this.createStoreModalClose()
        },
        async updateStore () {
            await this.$axios.put('/api/store/update', this.editStore)
            this.fetchStores()
            this.editStoreModalClose()
        },
        deleteStore (id, index) {
            console.log(id, index)
            this.$axios.delete(`/api/store/delete/${id}`)
            .then(res => {
                this.list = this.list.splice(index, 1)
            })
            .catch(err => {
                console.log(err)
            })
        },
        createStoreModalOpen () {
            this.isCreateStore = true
            this.newStore = {}
        },
        createStoreModalClose () {
            this.isCreateStore = false
            this.newStore = {}
        },
        editStoreModalOpen (store) {
            this.isEditStore = true
            this.editStore = store
        },
        editStoreModalClose () {
            this.isEditStore = false
            this.editStore = {}
        },
    }
}
</script>
