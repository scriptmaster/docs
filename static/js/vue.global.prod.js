/*!
 * Minimal Vue.js stub for documentation
 * Full Vue.js would be loaded from CDN or can be embedded separately
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
