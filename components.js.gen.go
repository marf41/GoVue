// Code generated by github.com/tv42/becky -- DO NOT EDIT.

package govue

var _ = IntWrap(asset{Name: "components.js", Content: "" +
	"GoVue = {\n    install(Vue, options) {\n    }\n}\n\nVue.use(GoVue)\nVue.use(Vuetify)\n\nVue.component('v-footer-cmp', { template: '#tmp-footer', data() { return { text: '' } },\n    methods: { set(s) { this.text = s } },\n    computed: { value() { var y = new Date().getFullYear(); return this.text || ('&copy; ' + y) } } }) \nVue.component('v-snack', { template: '#tmp-snack', data() { return { text: '', shw: false, color: '', timeout: 3000 } },\n    methods: { set(s, c) { this.text = s; this.color = c || ''; this.shw = true },\n                show() { this.shw = true } } })" +
	"", etag: `"fFtzA0mAz5o="`})