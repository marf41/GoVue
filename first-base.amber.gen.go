// Code generated by github.com/tv42/becky -- DO NOT EDIT.

package govue

var _ = IntWrap(asset{Name: "first-base.amber", Content: "" +
	"doctype html\nhtml[lang=Lang]\n  head\n    title #{$.Title}\n    meta[name=\"viewport\"]\n      [content=\"width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no, minimal-ui\"]\n    link[rel=\"stylesheet\"][href=\"/materialdesignicons.min.css\"]\n    link[rel=\"stylesheet\"][href=\"/vuetify.min.css\"]\n    link[rel=\"stylesheet\"][href=\"/roboto.css\"]\n    link[rel=\"stylesheet\"][href=\"/mdi.css\"]\n    link[rel=\"stylesheet\"][href=\"/MarqueeText.css\"]\n    each $url in $.Styles\n      link[rel=\"stylesheet\"][href=\"/\"+$url]\n    block head\n  body\n    block body\n    div#app\n      v-app\n        block 216672ef-bf07-400f-9eb7-9dcac1a2de0dapp\n        block app\n          | No app defined\n        v-snack[ref=\"snack\"]\n        v-footer-cmp[ref=\"footer\"]\n\n    script\n      var user = #{ $.User }\n    script[src=\"/vue.js\"]\n    script[src=\"/vuetify.js\"]\n\n    block 216672ef-bf07-400f-9eb7-9dcac1a2de0dcomponents\n    block components\n\n    each $url in $.Scripts\n      script[src=\"/\" + $url]\n\n    block scripts" +
	"", etag: `"wZloF+Pr5qI="`})
