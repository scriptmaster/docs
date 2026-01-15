/*!
 * Minimal Vue.js stub for documentation
 * This is a lightweight stub providing basic Vue.js API compatibility.
 * For full Vue.js features, replace this file with the complete Vue.js library from https://cdn.jsdelivr.net/npm/vue@3/dist/vue.global.prod.js
 * 
 * Supported: createApp(), mount(), basic data/mounted lifecycle
 * Not supported: Reactivity, templates, directives, components, router, etc.
 */
const Vue = {
  createApp: function(config) {
    return {
      mount: function(selector) {
        console.log('Vue app mounted on', selector);
        return this;
      },
      data: config.data || function() { return {}; },
      mounted: config.mounted || function() {}
    };
  }
};
if (typeof window !== 'undefined') {
  window.Vue = Vue;
}
