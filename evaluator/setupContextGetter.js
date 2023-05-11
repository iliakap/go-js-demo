const handler = {
    get(target, key) {
	 			return getCtx(key);
    },
	   };
const $ = new Proxy({}, handler);