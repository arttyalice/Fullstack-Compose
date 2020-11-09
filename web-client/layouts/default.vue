<template>
  <div class="container">
    <CThemeProvider>
      <CColorModeProvider>
        <CBox font-family="body" as="main">
          <c-menu>
            <c-menu-button
              variant-color="blue"
              size="sm"
              style="position: absolute; top: 10px; left: 10px;"
            >
              ☰
            </c-menu-button>
            <c-menu-list>
              <c-menu-item as="a" href="/store">
                Store
              </c-menu-item>
              <c-menu-item as="a" href="/product">
                Product
              </c-menu-item>
            </c-menu-list>
          </c-menu>
          <CReset />
          <Nuxt />
        </CBox>
      </CColorModeProvider>
    </CThemeProvider>
  </div>
</template>
<script>
import {
  CThemeProvider,
  CColorModeProvider,
  CReset,
  CBox
} from '@chakra-ui/vue'

export default {
  name: 'App',
  components: {
    CThemeProvider,
    CColorModeProvider,
    CReset,
    CBox
  }
}
</script>
