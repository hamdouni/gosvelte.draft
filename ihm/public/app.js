
(function(l, r) { if (l.getElementById('livereloadscript')) return; r = l.createElement('script'); r.async = 1; r.src = '//' + (window.location.host || 'localhost').split(':')[0] + ':35729/livereload.js?snipver=1'; r.id = 'livereloadscript'; l.getElementsByTagName('head')[0].appendChild(r) })(window.document);
var app = (function () {
    'use strict';

    function noop() { }
    function add_location(element, file, line, column, char) {
        element.__svelte_meta = {
            loc: { file, line, column, char }
        };
    }
    function run(fn) {
        return fn();
    }
    function blank_object() {
        return Object.create(null);
    }
    function run_all(fns) {
        fns.forEach(run);
    }
    function is_function(thing) {
        return typeof thing === 'function';
    }
    function safe_not_equal(a, b) {
        return a != a ? b == b : a !== b || ((a && typeof a === 'object') || typeof a === 'function');
    }
    function is_empty(obj) {
        return Object.keys(obj).length === 0;
    }

    function append(target, node) {
        target.appendChild(node);
    }
    function insert(target, node, anchor) {
        target.insertBefore(node, anchor || null);
    }
    function detach(node) {
        node.parentNode.removeChild(node);
    }
    function destroy_each(iterations, detaching) {
        for (let i = 0; i < iterations.length; i += 1) {
            if (iterations[i])
                iterations[i].d(detaching);
        }
    }
    function element(name) {
        return document.createElement(name);
    }
    function text(data) {
        return document.createTextNode(data);
    }
    function space() {
        return text(' ');
    }
    function listen(node, event, handler, options) {
        node.addEventListener(event, handler, options);
        return () => node.removeEventListener(event, handler, options);
    }
    function attr(node, attribute, value) {
        if (value == null)
            node.removeAttribute(attribute);
        else if (node.getAttribute(attribute) !== value)
            node.setAttribute(attribute, value);
    }
    function children(element) {
        return Array.from(element.childNodes);
    }
    function set_input_value(input, value) {
        input.value = value == null ? '' : value;
    }
    function custom_event(type, detail) {
        const e = document.createEvent('CustomEvent');
        e.initCustomEvent(type, false, false, detail);
        return e;
    }

    let current_component;
    function set_current_component(component) {
        current_component = component;
    }

    const dirty_components = [];
    const binding_callbacks = [];
    const render_callbacks = [];
    const flush_callbacks = [];
    const resolved_promise = Promise.resolve();
    let update_scheduled = false;
    function schedule_update() {
        if (!update_scheduled) {
            update_scheduled = true;
            resolved_promise.then(flush);
        }
    }
    function add_render_callback(fn) {
        render_callbacks.push(fn);
    }
    let flushing = false;
    const seen_callbacks = new Set();
    function flush() {
        if (flushing)
            return;
        flushing = true;
        do {
            // first, call beforeUpdate functions
            // and update components
            for (let i = 0; i < dirty_components.length; i += 1) {
                const component = dirty_components[i];
                set_current_component(component);
                update(component.$$);
            }
            set_current_component(null);
            dirty_components.length = 0;
            while (binding_callbacks.length)
                binding_callbacks.pop()();
            // then, once components are updated, call
            // afterUpdate functions. This may cause
            // subsequent updates...
            for (let i = 0; i < render_callbacks.length; i += 1) {
                const callback = render_callbacks[i];
                if (!seen_callbacks.has(callback)) {
                    // ...so guard against infinite loops
                    seen_callbacks.add(callback);
                    callback();
                }
            }
            render_callbacks.length = 0;
        } while (dirty_components.length);
        while (flush_callbacks.length) {
            flush_callbacks.pop()();
        }
        update_scheduled = false;
        flushing = false;
        seen_callbacks.clear();
    }
    function update($$) {
        if ($$.fragment !== null) {
            $$.update();
            run_all($$.before_update);
            const dirty = $$.dirty;
            $$.dirty = [-1];
            $$.fragment && $$.fragment.p($$.ctx, dirty);
            $$.after_update.forEach(add_render_callback);
        }
    }
    const outroing = new Set();
    let outros;
    function transition_in(block, local) {
        if (block && block.i) {
            outroing.delete(block);
            block.i(local);
        }
    }
    function transition_out(block, local, detach, callback) {
        if (block && block.o) {
            if (outroing.has(block))
                return;
            outroing.add(block);
            outros.c.push(() => {
                outroing.delete(block);
                if (callback) {
                    if (detach)
                        block.d(1);
                    callback();
                }
            });
            block.o(local);
        }
    }

    const globals = (typeof window !== 'undefined'
        ? window
        : typeof globalThis !== 'undefined'
            ? globalThis
            : global);
    function create_component(block) {
        block && block.c();
    }
    function mount_component(component, target, anchor, customElement) {
        const { fragment, on_mount, on_destroy, after_update } = component.$$;
        fragment && fragment.m(target, anchor);
        if (!customElement) {
            // onMount happens before the initial afterUpdate
            add_render_callback(() => {
                const new_on_destroy = on_mount.map(run).filter(is_function);
                if (on_destroy) {
                    on_destroy.push(...new_on_destroy);
                }
                else {
                    // Edge case - component was destroyed immediately,
                    // most likely as a result of a binding initialising
                    run_all(new_on_destroy);
                }
                component.$$.on_mount = [];
            });
        }
        after_update.forEach(add_render_callback);
    }
    function destroy_component(component, detaching) {
        const $$ = component.$$;
        if ($$.fragment !== null) {
            run_all($$.on_destroy);
            $$.fragment && $$.fragment.d(detaching);
            // TODO null out other refs, including component.$$ (but need to
            // preserve final state?)
            $$.on_destroy = $$.fragment = null;
            $$.ctx = [];
        }
    }
    function make_dirty(component, i) {
        if (component.$$.dirty[0] === -1) {
            dirty_components.push(component);
            schedule_update();
            component.$$.dirty.fill(0);
        }
        component.$$.dirty[(i / 31) | 0] |= (1 << (i % 31));
    }
    function init(component, options, instance, create_fragment, not_equal, props, dirty = [-1]) {
        const parent_component = current_component;
        set_current_component(component);
        const $$ = component.$$ = {
            fragment: null,
            ctx: null,
            // state
            props,
            update: noop,
            not_equal,
            bound: blank_object(),
            // lifecycle
            on_mount: [],
            on_destroy: [],
            on_disconnect: [],
            before_update: [],
            after_update: [],
            context: new Map(parent_component ? parent_component.$$.context : options.context || []),
            // everything else
            callbacks: blank_object(),
            dirty,
            skip_bound: false
        };
        let ready = false;
        $$.ctx = instance
            ? instance(component, options.props || {}, (i, ret, ...rest) => {
                const value = rest.length ? rest[0] : ret;
                if ($$.ctx && not_equal($$.ctx[i], $$.ctx[i] = value)) {
                    if (!$$.skip_bound && $$.bound[i])
                        $$.bound[i](value);
                    if (ready)
                        make_dirty(component, i);
                }
                return ret;
            })
            : [];
        $$.update();
        ready = true;
        run_all($$.before_update);
        // `false` as a special case of no DOM component
        $$.fragment = create_fragment ? create_fragment($$.ctx) : false;
        if (options.target) {
            if (options.hydrate) {
                const nodes = children(options.target);
                // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
                $$.fragment && $$.fragment.l(nodes);
                nodes.forEach(detach);
            }
            else {
                // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
                $$.fragment && $$.fragment.c();
            }
            if (options.intro)
                transition_in(component.$$.fragment);
            mount_component(component, options.target, options.anchor, options.customElement);
            flush();
        }
        set_current_component(parent_component);
    }
    /**
     * Base class for Svelte components. Used when dev=false.
     */
    class SvelteComponent {
        $destroy() {
            destroy_component(this, 1);
            this.$destroy = noop;
        }
        $on(type, callback) {
            const callbacks = (this.$$.callbacks[type] || (this.$$.callbacks[type] = []));
            callbacks.push(callback);
            return () => {
                const index = callbacks.indexOf(callback);
                if (index !== -1)
                    callbacks.splice(index, 1);
            };
        }
        $set($$props) {
            if (this.$$set && !is_empty($$props)) {
                this.$$.skip_bound = true;
                this.$$set($$props);
                this.$$.skip_bound = false;
            }
        }
    }

    function dispatch_dev(type, detail) {
        document.dispatchEvent(custom_event(type, Object.assign({ version: '3.37.0' }, detail)));
    }
    function append_dev(target, node) {
        dispatch_dev('SvelteDOMInsert', { target, node });
        append(target, node);
    }
    function insert_dev(target, node, anchor) {
        dispatch_dev('SvelteDOMInsert', { target, node, anchor });
        insert(target, node, anchor);
    }
    function detach_dev(node) {
        dispatch_dev('SvelteDOMRemove', { node });
        detach(node);
    }
    function listen_dev(node, event, handler, options, has_prevent_default, has_stop_propagation) {
        const modifiers = options === true ? ['capture'] : options ? Array.from(Object.keys(options)) : [];
        if (has_prevent_default)
            modifiers.push('preventDefault');
        if (has_stop_propagation)
            modifiers.push('stopPropagation');
        dispatch_dev('SvelteDOMAddEventListener', { node, event, handler, modifiers });
        const dispose = listen(node, event, handler, options);
        return () => {
            dispatch_dev('SvelteDOMRemoveEventListener', { node, event, handler, modifiers });
            dispose();
        };
    }
    function attr_dev(node, attribute, value) {
        attr(node, attribute, value);
        if (value == null)
            dispatch_dev('SvelteDOMRemoveAttribute', { node, attribute });
        else
            dispatch_dev('SvelteDOMSetAttribute', { node, attribute, value });
    }
    function set_data_dev(text, data) {
        data = '' + data;
        if (text.wholeText === data)
            return;
        dispatch_dev('SvelteDOMSetData', { node: text, data });
        text.data = data;
    }
    function validate_each_argument(arg) {
        if (typeof arg !== 'string' && !(arg && typeof arg === 'object' && 'length' in arg)) {
            let msg = '{#each} only iterates over array-like objects.';
            if (typeof Symbol === 'function' && arg && Symbol.iterator in arg) {
                msg += ' You can use a spread to convert this iterable into an array.';
            }
            throw new Error(msg);
        }
    }
    function validate_slots(name, slot, keys) {
        for (const slot_key of Object.keys(slot)) {
            if (!~keys.indexOf(slot_key)) {
                console.warn(`<${name}> received an unexpected slot "${slot_key}".`);
            }
        }
    }
    /**
     * Base class for Svelte components with some minor dev-enhancements. Used when dev=true.
     */
    class SvelteComponentDev extends SvelteComponent {
        constructor(options) {
            if (!options || (!options.target && !options.$$inline)) {
                throw new Error("'target' is a required option");
            }
            super();
        }
        $destroy() {
            super.$destroy();
            this.$destroy = () => {
                console.warn('Component was already destroyed'); // eslint-disable-line no-console
            };
        }
        $capture_state() { }
        $inject_state() { }
    }

    /* ihm/Login.svelte generated by Svelte v3.37.0 */

    const file$1 = "ihm/Login.svelte";

    function create_fragment$1(ctx) {
    	let form;
    	let input0;
    	let t0;
    	let input1;
    	let t1;
    	let button;

    	const block = {
    		c: function create() {
    			form = element("form");
    			input0 = element("input");
    			t0 = space();
    			input1 = element("input");
    			t1 = space();
    			button = element("button");
    			button.textContent = "Connecter";
    			attr_dev(input0, "type", "text");
    			attr_dev(input0, "name", "username");
    			attr_dev(input0, "placeholder", "Utilisateur");
    			add_location(input0, file$1, 1, 4, 41);
    			attr_dev(input1, "type", "password");
    			attr_dev(input1, "name", "password");
    			attr_dev(input1, "placeholder", "Mot de passe");
    			add_location(input1, file$1, 2, 4, 107);
    			attr_dev(button, "type", "submit");
    			add_location(button, file$1, 3, 4, 178);
    			attr_dev(form, "action", "/login");
    			attr_dev(form, "method", "post");
    			add_location(form, file$1, 0, 0, 0);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, form, anchor);
    			append_dev(form, input0);
    			append_dev(form, t0);
    			append_dev(form, input1);
    			append_dev(form, t1);
    			append_dev(form, button);
    		},
    		p: noop,
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(form);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$1.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$1($$self, $$props) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots("Login", slots, []);
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== "$$") console.warn(`<Login> was created with unknown prop '${key}'`);
    	});

    	return [];
    }

    class Login extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$1, create_fragment$1, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Login",
    			options,
    			id: create_fragment$1.name
    		});
    	}
    }

    /* ihm/App.svelte generated by Svelte v3.37.0 */

    const { console: console_1 } = globals;
    const file = "ihm/App.svelte";

    function get_each_context(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[10] = list[i];
    	return child_ctx;
    }

    // (70:4) {#each historique as evenement}
    function create_each_block(ctx) {
    	let li;
    	let t_value = /*evenement*/ ctx[10] + "";
    	let t;

    	const block = {
    		c: function create() {
    			li = element("li");
    			t = text(t_value);
    			add_location(li, file, 70, 8, 1756);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, li, anchor);
    			append_dev(li, t);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*historique*/ 4 && t_value !== (t_value = /*evenement*/ ctx[10] + "")) set_data_dev(t, t_value);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(li);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block.name,
    		type: "each",
    		source: "(70:4) {#each historique as evenement}",
    		ctx
    	});

    	return block;
    }

    function create_fragment(ctx) {
    	let pre;
    	let t1;
    	let hr0;
    	let t2;
    	let input;
    	let t3;
    	let button0;
    	let t5;
    	let button1;
    	let t7;
    	let button2;
    	let t9;
    	let hr1;
    	let t10;
    	let h1;
    	let t11;
    	let t12;
    	let hr2;
    	let t13;
    	let ul;
    	let t14;
    	let login;
    	let t15;
    	let hr3;
    	let t16;
    	let a0;
    	let t18;
    	let hr4;
    	let t19;
    	let t20;
    	let t21;
    	let hr5;
    	let t22;
    	let a1;
    	let current;
    	let mounted;
    	let dispose;
    	let each_value = /*historique*/ ctx[2];
    	validate_each_argument(each_value);
    	let each_blocks = [];

    	for (let i = 0; i < each_value.length; i += 1) {
    		each_blocks[i] = create_each_block(get_each_context(ctx, each_value, i));
    	}

    	login = new Login({ $$inline: true });

    	const block = {
    		c: function create() {
    			pre = element("pre");
    			pre.textContent = "Bonjour depuis Svelte.";
    			t1 = space();
    			hr0 = element("hr");
    			t2 = space();
    			input = element("input");
    			t3 = space();
    			button0 = element("button");
    			button0.textContent = "Bonjour";
    			t5 = space();
    			button1 = element("button");
    			button1.textContent = "Majuscule";
    			t7 = space();
    			button2 = element("button");
    			button2.textContent = "Minuscule";
    			t9 = space();
    			hr1 = element("hr");
    			t10 = space();
    			h1 = element("h1");
    			t11 = text(/*resultat*/ ctx[1]);
    			t12 = space();
    			hr2 = element("hr");
    			t13 = space();
    			ul = element("ul");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			t14 = space();
    			create_component(login.$$.fragment);
    			t15 = space();
    			hr3 = element("hr");
    			t16 = space();
    			a0 = element("a");
    			a0.textContent = "Logout";
    			t18 = space();
    			hr4 = element("hr");
    			t19 = text("\nConnexion : ");
    			t20 = text(/*isConnected*/ ctx[3]);
    			t21 = space();
    			hr5 = element("hr");
    			t22 = space();
    			a1 = element("a");
    			a1.textContent = "Check";
    			attr_dev(pre, "class", "toto svelte-1tj96w7");
    			add_location(pre, file, 51, 0, 1379);
    			add_location(hr0, file, 54, 0, 1432);
    			attr_dev(input, "type", "text");
    			attr_dev(input, "placeholder", "Entrer votre prénom...");
    			add_location(input, file, 55, 0, 1437);
    			add_location(button0, file, 56, 0, 1513);
    			add_location(button1, file, 59, 0, 1569);
    			add_location(button2, file, 62, 0, 1623);
    			add_location(hr1, file, 65, 0, 1677);
    			add_location(h1, file, 66, 0, 1682);
    			add_location(hr2, file, 67, 0, 1702);
    			add_location(ul, file, 68, 0, 1707);
    			add_location(hr3, file, 74, 0, 1804);
    			attr_dev(a0, "href", "/logout");
    			add_location(a0, file, 75, 0, 1809);
    			add_location(hr4, file, 76, 0, 1838);
    			add_location(hr5, file, 78, 0, 1869);
    			attr_dev(a1, "href", "#check");
    			add_location(a1, file, 79, 0, 1874);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, pre, anchor);
    			insert_dev(target, t1, anchor);
    			insert_dev(target, hr0, anchor);
    			insert_dev(target, t2, anchor);
    			insert_dev(target, input, anchor);
    			set_input_value(input, /*nom*/ ctx[0]);
    			insert_dev(target, t3, anchor);
    			insert_dev(target, button0, anchor);
    			insert_dev(target, t5, anchor);
    			insert_dev(target, button1, anchor);
    			insert_dev(target, t7, anchor);
    			insert_dev(target, button2, anchor);
    			insert_dev(target, t9, anchor);
    			insert_dev(target, hr1, anchor);
    			insert_dev(target, t10, anchor);
    			insert_dev(target, h1, anchor);
    			append_dev(h1, t11);
    			insert_dev(target, t12, anchor);
    			insert_dev(target, hr2, anchor);
    			insert_dev(target, t13, anchor);
    			insert_dev(target, ul, anchor);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(ul, null);
    			}

    			insert_dev(target, t14, anchor);
    			mount_component(login, target, anchor);
    			insert_dev(target, t15, anchor);
    			insert_dev(target, hr3, anchor);
    			insert_dev(target, t16, anchor);
    			insert_dev(target, a0, anchor);
    			insert_dev(target, t18, anchor);
    			insert_dev(target, hr4, anchor);
    			insert_dev(target, t19, anchor);
    			insert_dev(target, t20, anchor);
    			insert_dev(target, t21, anchor);
    			insert_dev(target, hr5, anchor);
    			insert_dev(target, t22, anchor);
    			insert_dev(target, a1, anchor);
    			current = true;

    			if (!mounted) {
    				dispose = [
    					listen_dev(input, "input", /*input_input_handler*/ ctx[8]),
    					listen_dev(button0, "click", /*callBonjour*/ ctx[4], false, false, false),
    					listen_dev(button1, "click", /*callMaj*/ ctx[5], false, false, false),
    					listen_dev(button2, "click", /*callMin*/ ctx[6], false, false, false),
    					listen_dev(a1, "click", /*checkConnexion*/ ctx[7], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*nom*/ 1 && input.value !== /*nom*/ ctx[0]) {
    				set_input_value(input, /*nom*/ ctx[0]);
    			}

    			if (!current || dirty & /*resultat*/ 2) set_data_dev(t11, /*resultat*/ ctx[1]);

    			if (dirty & /*historique*/ 4) {
    				each_value = /*historique*/ ctx[2];
    				validate_each_argument(each_value);
    				let i;

    				for (i = 0; i < each_value.length; i += 1) {
    					const child_ctx = get_each_context(ctx, each_value, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    					} else {
    						each_blocks[i] = create_each_block(child_ctx);
    						each_blocks[i].c();
    						each_blocks[i].m(ul, null);
    					}
    				}

    				for (; i < each_blocks.length; i += 1) {
    					each_blocks[i].d(1);
    				}

    				each_blocks.length = each_value.length;
    			}

    			if (!current || dirty & /*isConnected*/ 8) set_data_dev(t20, /*isConnected*/ ctx[3]);
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(login.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(login.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(pre);
    			if (detaching) detach_dev(t1);
    			if (detaching) detach_dev(hr0);
    			if (detaching) detach_dev(t2);
    			if (detaching) detach_dev(input);
    			if (detaching) detach_dev(t3);
    			if (detaching) detach_dev(button0);
    			if (detaching) detach_dev(t5);
    			if (detaching) detach_dev(button1);
    			if (detaching) detach_dev(t7);
    			if (detaching) detach_dev(button2);
    			if (detaching) detach_dev(t9);
    			if (detaching) detach_dev(hr1);
    			if (detaching) detach_dev(t10);
    			if (detaching) detach_dev(h1);
    			if (detaching) detach_dev(t12);
    			if (detaching) detach_dev(hr2);
    			if (detaching) detach_dev(t13);
    			if (detaching) detach_dev(ul);
    			destroy_each(each_blocks, detaching);
    			if (detaching) detach_dev(t14);
    			destroy_component(login, detaching);
    			if (detaching) detach_dev(t15);
    			if (detaching) detach_dev(hr3);
    			if (detaching) detach_dev(t16);
    			if (detaching) detach_dev(a0);
    			if (detaching) detach_dev(t18);
    			if (detaching) detach_dev(hr4);
    			if (detaching) detach_dev(t19);
    			if (detaching) detach_dev(t20);
    			if (detaching) detach_dev(t21);
    			if (detaching) detach_dev(hr5);
    			if (detaching) detach_dev(t22);
    			if (detaching) detach_dev(a1);
    			mounted = false;
    			run_all(dispose);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    async function callApi(endpoint) {
    	let url = endpoint;

    	try {
    		let response = await fetch(url);

    		if (response.ok) {
    			return response.json();
    		} else {
    			console.log("Erreur http " + response.status + " sur url " + url);
    		}
    	} catch(error) {
    		console.log("Erreur réseau " + error);
    	}
    }

    function instance($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots("App", slots, []);
    	let nom = "";
    	let resultat = "";
    	let historique = [];
    	let isConnected = false;

    	async function callBonjour() {
    		let res = await callApi("/hello?nom=" + nom);
    		if (res != null) $$invalidate(1, resultat = res);
    		reloadHistoric();
    	}

    	async function callMaj() {
    		let res = await callApi("/upper?nom=" + nom);
    		if (res != null) $$invalidate(1, resultat = res);
    		reloadHistoric();
    	}

    	async function callMin() {
    		let res = await callApi("/lower?nom=" + nom);
    		if (res != null) $$invalidate(1, resultat = res);
    		reloadHistoric();
    	}

    	async function reloadHistoric() {
    		let res = await callApi("/historic");
    		if (res != null) $$invalidate(2, historique = res);
    	}

    	reloadHistoric();

    	async function checkConnexion() {
    		let response = await fetch("check");

    		if (response.ok) {
    			$$invalidate(3, isConnected = true);
    			console.log("HOURA");
    		}
    	}

    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== "$$") console_1.warn(`<App> was created with unknown prop '${key}'`);
    	});

    	function input_input_handler() {
    		nom = this.value;
    		$$invalidate(0, nom);
    	}

    	$$self.$capture_state = () => ({
    		Login,
    		nom,
    		resultat,
    		historique,
    		isConnected,
    		callApi,
    		callBonjour,
    		callMaj,
    		callMin,
    		reloadHistoric,
    		checkConnexion
    	});

    	$$self.$inject_state = $$props => {
    		if ("nom" in $$props) $$invalidate(0, nom = $$props.nom);
    		if ("resultat" in $$props) $$invalidate(1, resultat = $$props.resultat);
    		if ("historique" in $$props) $$invalidate(2, historique = $$props.historique);
    		if ("isConnected" in $$props) $$invalidate(3, isConnected = $$props.isConnected);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [
    		nom,
    		resultat,
    		historique,
    		isConnected,
    		callBonjour,
    		callMaj,
    		callMin,
    		checkConnexion,
    		input_input_handler
    	];
    }

    class App extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance, create_fragment, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "App",
    			options,
    			id: create_fragment.name
    		});
    	}
    }

    const app = new App({
    	target: document.body,
    	props: {}
    });

    return app;

}());
//# sourceMappingURL=app.js.map
