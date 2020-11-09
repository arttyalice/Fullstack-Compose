<template>
  <div class="container">
    <c-box
        d="flex"
        justify-content="flex-end"
        p="2"
    >
        <c-button @click="productModalOpen">Add Product</c-button>
    </c-box>
    <c-box
        p="5"
    >
        <c-simple-grid :columns="3" :spacing="2">
            <c-box
                p="3"
                border-width="1px"
                justify-content="center"
                v-for="(product, index) in list"
                :key="product.id"
            >
                <c-heading text-align="center" size="xl">
                    {{ product.name }}
                </c-heading>
                <c-flex justify="center" direction="column" align="center">
                    Description: {{ product.description }}<br />
                    Price: {{ product.price_per_unit }}<br />
                    Unit: {{ product.unit_name }}<br />
                    Stock: {{ product.quantity }}
                </c-flex>
                <c-flex justify="center" direction="column">
                    <c-button as="a" @click="editProductModalOpen(product)">
                        Edit
                    </c-button>
                    <c-button variant-color="red" @click="deleteProduct(product.id, index)">
                        Delete
                    </c-button>
                </c-flex>
            </c-box>
        </c-simple-grid>
    </c-box>

    <!-- Create -->
    <c-modal
      :is-open="isCreateProduct"
      :on-close="productModalClose"
    >
      <c-modal-content>
        <c-modal-header>Create Product</c-modal-header>
        <c-modal-close-button />
        <c-modal-body>
          <c-form-control is-required>
            <c-form-label>Category</c-form-label>
            <c-select placeholder="Select Category" v-model="newProduct['category_id']">
                <option
                    v-for="category in categorys"
                    :value="category.id"
                    :key="category.id"
                >
                    {{ category.name }}
                </option>
            </c-select>
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Name</c-form-label>
            <c-input v-model="newProduct['name']" placeholder="Product Name" />
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Price</c-form-label>
            <c-input v-model="newProduct['price_per_unit']" placeholder="Product Price" />
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Unit name</c-form-label>
            <c-input v-model="newProduct['unit_name']" placeholder="Product Unit name" />
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Quantity</c-form-label>
            <c-input v-model="newProduct['quantity']" placeholder="Product Quantity" />
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Description</c-form-label>
            <c-textarea v-model="newProduct['description']" placeholder="Product Description" />
          </c-form-control>
        </c-modal-body>
        <c-modal-footer>
          <c-button @click="productModalClose" variant-color="blue" mr="3">
            Cancel
          </c-button>
          <c-button @click="createProduct">Save</c-button>
        </c-modal-footer>
      </c-modal-content>
      <c-modal-overlay />
    </c-modal>

    <!-- Edit -->
    <c-modal
      :is-open="isEditProduct"
      :on-close="editProductModalClose"
    >
      <c-modal-content>
        <c-modal-header>Edit Product</c-modal-header>
        <c-modal-close-button />
        <c-modal-body>
          <c-form-control is-required>
            <c-form-label>Category</c-form-label>
            <c-select placeholder="Select Category" v-model="editProduct['category_id']">
                <option
                    v-for="category in categorys"
                    :value="category.id"
                    :key="category.id"
                >
                    {{ category.name }}
                </option>
            </c-select>
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Name</c-form-label>
            <c-input v-model="editProduct['name']" placeholder="Product Name" />
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Price</c-form-label>
            <c-input v-model="editProduct['price_per_unit']" placeholder="Product Price" />
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Unit name</c-form-label>
            <c-input v-model="editProduct['unit_name']" placeholder="Product Unit name" />
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Quantity</c-form-label>
            <c-input v-model="editProduct['quantity']" placeholder="Product Quantity" />
          </c-form-control>

          <c-form-control mt="4" is-required>
            <c-form-label>Description</c-form-label>
            <c-textarea v-model="editProduct['description']" placeholder="Product Description" />
          </c-form-control>
        </c-modal-body>
        <c-modal-footer>
          <c-button @click="editProductModalClose" variant-color="blue" mr="3">
            Cancel
          </c-button>
          <c-button @click="createProduct">Save</c-button>
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
            categorys: [],
            // Create product data
            isCreateProduct: false,
            newProduct: {},
            // Edit product data
            // Create product data
            isEditProduct: false,
            editProduct: {},
		}
    },
    created () {
        this.fetchProducts()
        this.fetchCategory()
    },
    methods: {
        fetchProducts () {
            this.$axios.get("/api/product/list")
            .then(res => {
                this.list = res.data.items
            })
            .catch(err => {
                console.log(err)
            })
        },
        fetchCategory () {
            this.$axios.get("/api/product/category/list")
            .then(res => {
                this.categorys = res.data.items
            })
            .catch(err => {
                console.log(err)
            })
        },
        deleteProduct (id, index) {
            this.$axios.delete(`/api/product/delete/${id}`)
            .then(res => {
                this.list = this.list.filter(item => item.id !== id)
            })
            .catch(err => {
                console.log(err)
            })
        },
        async createProduct () {
            if (!this.newProduct.name ||
                !this.newProduct.description ||
                !this.newProduct.price_per_unit ||
                !this.newProduct.unit_name ||
                !this.newProduct.category_id ||
                !this.newProduct.quantity
            ) {
                this.$toast({
                    title: 'Please complete the form.',
                    description: "Please fill all required fields in the form.",
                    status: 'error',
                    duration: 10000
                })
                return
            }
            await this.$axios.post(`/api/product/create`, {
                name: this.newProduct.name,
                description: this.newProduct.description,
                price_per_unit: Number(this.newProduct.price_per_unit),
                unit_name: this.newProduct.unit_name,
                category_id: Number(this.newProduct.category_id),
                quantity: Number(this.newProduct.quantity),
            })
            this.productModalClose()
            this.fetchProducts()
        },
        async updateProduct () {

        },
        productModalOpen () {
            this.isCreateProduct = true
            this.newProduct = {}
        },
        productModalClose () {
            this.isCreateProduct = false
            this.newProduct = {}
        },
        editProductModalOpen (product) {
            console.log(product)
            this.isEditProduct = true
            this.editProduct = product
        },
        editProductModalClose () {
            this.isEditProduct = false
            this.editProduct = {}
        },
    }
}
</script>
