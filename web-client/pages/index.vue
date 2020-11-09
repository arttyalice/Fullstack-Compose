<template>
  <div class="container">
    <c-box
      v-bind="mainStyles[colorMode]"
      d="flex"
      w="100vw"
      h="100vh"
      flex-dir="column"
      justify-content="center"
    >
      <c-heading text-align="center" mb="4"> Hello Manager! </c-heading>
      <c-flex justify="center" direction="column" align="center">
		<c-stack :spacing="5" style="width: 300px;">
			<c-box :p="5" shadow="sm" border-width="0px" as="a" href="/store">
				<c-heading>Store</c-heading>
				<c-text :mt="4">Manage Store.</c-text>
			</c-box>
			<c-box :p="5" shadow="sm" border-width="0px" as="a" href="/product">
				<c-heading>Product</c-heading>
				<c-text :mt="4">Manage Product</c-text>
			</c-box>
		</c-stack>
      </c-flex>
    </c-box>
  </div>
</template>

<script lang="js">

export default {
	name: 'App',
	inject: ['$chakraColorMode', '$toggleColorMode'],
	data () {
		return {
			showModal: false,
			mainStyles: {
				dark: {
					bg: 'gray.700',
					color: 'whiteAlpha.900'
				},
				light: {
					bg: 'white',
					color: 'gray.900'
				}
			}
		}
	},
	computed: {
		colorMode () {
			return this.$chakraColorMode()
		},
		theme () {
			return this.$chakraTheme()
		},
		toggleColorMode () {
			return this.$toggleColorMode
		}
	}
}
</script>
