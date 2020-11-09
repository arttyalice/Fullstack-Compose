<template>
  <div class="container">
    <c-box
        d="flex"
        justify-content="flex-end"
        p="2"
    >
        <c-button as="a" href="/store">Go back</c-button>
    </c-box>
    <c-box
        p="5"
    >
        <c-box
            p="3"
            border-width="1px"
            justify-content="center"
        >
            <c-heading size="xl">
                {{ info.name }}
            </c-heading>
            <c-flex>
                Description: {{ info.description }}<br />
                Phone: {{ info.phone_number }}<br />
                Address: {{ info.address }}
            </c-flex>
        </c-box>

        <c-box
            d="flex"
            justify-content="flex-end"
            p="2"
        >
            <c-button @click="productModalOpen">Add Product</c-button>
        </c-box>
        <c-simple-grid :columns="3" :spacing="2">
            <c-box
                d="flex"
                p="3"
                border-width="1px"
                justify-content="center"
                v-for="(product) in info.products"
                :key="product.id"
            >
                <c-heading size="xl">
                    {{ product.name }}
                </c-heading>
                <c-flex justify="center" direction="column" align="center">
                    Description: {{ product.description }}<br />
                    Price: {{ product.price_per_unit }}/{{ product.unit_name }}<br />
                    จำนวน: {{ product.quantity }}
                </c-flex>
                <c-flex justify="center" direction="column">
                    <c-button as="a" :href="`/product/${product.id}`">
                        info
                    </c-button>
                    <c-button variant-color="red" @click="deleteProduct(product.id)">
                        Delete
                    </c-button>
                </c-flex>
            </c-box>
        </c-simple-grid>
    </c-box>

    <c-modal
      :is-open="isAddProduct"
      :on-close="productModalClose"
    >
      <c-modal-content>
        <c-modal-header>Add Product</c-modal-header>
        <c-modal-close-button />
        <c-modal-body>
          <c-form-control is-required>
            <c-form-label>Product</c-form-label>
            <c-select placeholder="Select Product" v-model="add_products[0]['product_id']">
                <option
                    v-for="product in products"
                    :value="product.id"
                    :key="product.id"
                >
                    {{ product.name }}
                </option>
            </c-select>
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Quantity</c-form-label>
            <c-input v-model="add_products[0]['product_quantity']" placeholder="Product Quantity" />
          </c-form-control>
        </c-modal-body>
        <c-modal-footer>
          <c-button @click="productModalClose" variant-color="blue" mr="3">
            Cancel
          </c-button>
          <c-button @click="addProduct">Save</c-button>
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
                info: {},
                products: [],
                // add product data
                isAddProduct: false,
                add_products: [{
                    "store_id": null,
                    "product_id": null,
                    "product_quantity": null
                }]
            }
        },
        created () {
            this.fetchStore()
            this.fetchProduct()
        },
        methods: {
            fetchStore () {
                this.$axios.get(`/api/store/info/${this.$route.params.id}`)
                .then(res => {
                    this.info = res.data.item
                })
                .catch(err => {
                    console.log(err)
                })
            },
            fetchProduct () {
                this.$axios.get(`/api/product/list`)
                .then(res => {
                    this.products = res.data.items
                })
                .catch(err => {
                    console.log(err)
                })
            },
            async deleteProduct (id) {
                console.log(id)
                await this.$axios.delete(`/api/store/product/delete/${id}`)
                this.fetchStore()
            },
            async addProduct () {
                if (this.add_products[0]['product_id'] == null) {
                    this.$toast({
                        title: 'Please Select Product.',
                        description: "Please select product and fill quantity before submit.",
                        status: 'error',
                        duration: 10000
                    })
                    return
                }
                if (this.add_products[0]['product_quantity'] == null) {
                    this.add_products[0]['product_quantity'] = 0
                }
                await this.$axios.post(`/api/store/product/create`, [{
                    product_quantity: Number(this.add_products[0]['product_quantity']),
                    product_id: Number(this.add_products[0]['product_id']),
                    store_id: Number(this.add_products[0]['store_id'])
                }])
                this.isAddProduct = false
                this.fetchStore()
            },
            productModalOpen () {
                this.isAddProduct = true
                this.add_products = [{
                    "store_id": this.$route.params.id,
                    "product_id": null,
                    "product_quantity": null
                }]
            },
            productModalClose () {
                this.isAddProduct = false
                this.add_products = [{
                    "store_id": this.$route.params.id,
                    "product_id": null,
                    "product_quantity": null
                }]
            }
        }
    }
</script>
