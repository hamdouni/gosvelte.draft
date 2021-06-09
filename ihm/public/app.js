
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
    function empty() {
        return text('');
    }
    function listen(node, event, handler, options) {
        node.addEventListener(event, handler, options);
        return () => node.removeEventListener(event, handler, options);
    }
    function prevent_default(fn) {
        return function (event) {
            event.preventDefault();
            // @ts-ignore
            return fn.call(this, event);
        };
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
    function toggle_class(element, name, toggle) {
        element.classList[toggle ? 'add' : 'remove'](name);
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
    function add_flush_callback(fn) {
        flush_callbacks.push(fn);
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
    function group_outros() {
        outros = {
            r: 0,
            c: [],
            p: outros // parent group
        };
    }
    function check_outros() {
        if (!outros.r) {
            run_all(outros.c);
        }
        outros = outros.p;
    }
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

    function destroy_block(block, lookup) {
        block.d(1);
        lookup.delete(block.key);
    }
    function update_keyed_each(old_blocks, dirty, get_key, dynamic, ctx, list, lookup, node, destroy, create_each_block, next, get_context) {
        let o = old_blocks.length;
        let n = list.length;
        let i = o;
        const old_indexes = {};
        while (i--)
            old_indexes[old_blocks[i].key] = i;
        const new_blocks = [];
        const new_lookup = new Map();
        const deltas = new Map();
        i = n;
        while (i--) {
            const child_ctx = get_context(ctx, list, i);
            const key = get_key(child_ctx);
            let block = lookup.get(key);
            if (!block) {
                block = create_each_block(key, child_ctx);
                block.c();
            }
            else if (dynamic) {
                block.p(child_ctx, dirty);
            }
            new_lookup.set(key, new_blocks[i] = block);
            if (key in old_indexes)
                deltas.set(key, Math.abs(i - old_indexes[key]));
        }
        const will_move = new Set();
        const did_move = new Set();
        function insert(block) {
            transition_in(block, 1);
            block.m(node, next);
            lookup.set(block.key, block);
            next = block.first;
            n--;
        }
        while (o && n) {
            const new_block = new_blocks[n - 1];
            const old_block = old_blocks[o - 1];
            const new_key = new_block.key;
            const old_key = old_block.key;
            if (new_block === old_block) {
                // do nothing
                next = new_block.first;
                o--;
                n--;
            }
            else if (!new_lookup.has(old_key)) {
                // remove old block
                destroy(old_block, lookup);
                o--;
            }
            else if (!lookup.has(new_key) || will_move.has(new_key)) {
                insert(new_block);
            }
            else if (did_move.has(old_key)) {
                o--;
            }
            else if (deltas.get(new_key) > deltas.get(old_key)) {
                did_move.add(new_key);
                insert(new_block);
            }
            else {
                will_move.add(old_key);
                o--;
            }
        }
        while (o--) {
            const old_block = old_blocks[o];
            if (!new_lookup.has(old_block.key))
                destroy(old_block, lookup);
        }
        while (n)
            insert(new_blocks[n - 1]);
        return new_blocks;
    }
    function validate_each_keys(ctx, list, get_context, get_key) {
        const keys = new Set();
        for (let i = 0; i < list.length; i++) {
            const key = get_key(get_context(ctx, list, i));
            if (keys.has(key)) {
                throw new Error('Cannot have duplicate keys in a keyed each');
            }
            keys.add(key);
        }
    }

    function bind(component, name, callback) {
        const index = component.$$.props[name];
        if (index !== undefined) {
            component.$$.bound[index] = callback;
            callback(component.$$.ctx[index]);
        }
    }
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

    /* Login.svelte generated by Svelte v3.37.0 */

    const file$5 = "Login.svelte";

    // (52:12) {#if err}
    function create_if_block$4(ctx) {
    	let article;
    	let div;
    	let strong;
    	let t;

    	const block = {
    		c: function create() {
    			article = element("article");
    			div = element("div");
    			strong = element("strong");
    			t = text(/*err*/ ctx[2]);
    			add_location(strong, file$5, 54, 20, 1966);
    			attr_dev(div, "class", "message-body");
    			add_location(div, file$5, 53, 16, 1919);
    			attr_dev(article, "class", "message is-danger");
    			add_location(article, file$5, 52, 12, 1867);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, article, anchor);
    			append_dev(article, div);
    			append_dev(div, strong);
    			append_dev(strong, t);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*err*/ 4) set_data_dev(t, /*err*/ ctx[2]);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(article);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$4.name,
    		type: "if",
    		source: "(52:12) {#if err}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$6(ctx) {
    	let login_1;
    	let form;
    	let div1;
    	let label0;
    	let t1;
    	let div0;
    	let input0;
    	let t2;
    	let span0;
    	let i0;
    	let t3;
    	let div3;
    	let label1;
    	let t5;
    	let div2;
    	let input1;
    	let t6;
    	let span1;
    	let i1;
    	let t7;
    	let div4;
    	let button;
    	let strong;
    	let t9;
    	let div5;
    	let mounted;
    	let dispose;
    	let if_block = /*err*/ ctx[2] && create_if_block$4(ctx);

    	const block = {
    		c: function create() {
    			login_1 = element("login");
    			form = element("form");
    			div1 = element("div");
    			label0 = element("label");
    			label0.textContent = "Identifiant";
    			t1 = space();
    			div0 = element("div");
    			input0 = element("input");
    			t2 = space();
    			span0 = element("span");
    			i0 = element("i");
    			t3 = space();
    			div3 = element("div");
    			label1 = element("label");
    			label1.textContent = "Mot de passe";
    			t5 = space();
    			div2 = element("div");
    			input1 = element("input");
    			t6 = space();
    			span1 = element("span");
    			i1 = element("i");
    			t7 = space();
    			div4 = element("div");
    			button = element("button");
    			strong = element("strong");
    			strong.textContent = "Connecter";
    			t9 = space();
    			div5 = element("div");
    			if (if_block) if_block.c();
    			attr_dev(label0, "for", "username");
    			attr_dev(label0, "class", "label");
    			add_location(label0, file$5, 26, 12, 786);
    			attr_dev(input0, "class", "input");
    			attr_dev(input0, "type", "text");
    			attr_dev(input0, "name", "username");
    			input0.autofocus = true;
    			add_location(input0, file$5, 28, 16, 923);
    			attr_dev(i0, "class", "fas fa-user");
    			add_location(i0, file$5, 30, 20, 1078);
    			attr_dev(span0, "class", "icon is-small is-left");
    			add_location(span0, file$5, 29, 16, 1021);
    			attr_dev(div0, "class", "control has-icons-left has-icons-right");
    			add_location(div0, file$5, 27, 12, 854);
    			attr_dev(div1, "class", "field");
    			add_location(div1, file$5, 25, 8, 754);
    			attr_dev(label1, "for", "password");
    			attr_dev(label1, "class", "label");
    			add_location(label1, file$5, 35, 12, 1204);
    			attr_dev(input1, "class", "input");
    			attr_dev(input1, "type", "password");
    			attr_dev(input1, "name", "password");
    			add_location(input1, file$5, 37, 16, 1342);
    			attr_dev(i1, "class", "fas fa-unlock-alt");
    			add_location(i1, file$5, 39, 20, 1491);
    			attr_dev(span1, "class", "icon is-small is-left");
    			add_location(span1, file$5, 38, 16, 1434);
    			attr_dev(div2, "class", "control has-icons-left has-icons-right");
    			add_location(div2, file$5, 36, 12, 1273);
    			attr_dev(div3, "class", "field");
    			add_location(div3, file$5, 34, 8, 1172);
    			add_location(strong, file$5, 45, 16, 1717);
    			attr_dev(button, "class", "button is-primary is-large");
    			add_location(button, file$5, 44, 12, 1625);
    			attr_dev(div4, "class", "buttons");
    			add_location(div4, file$5, 43, 8, 1591);
    			add_location(div5, file$5, 50, 8, 1827);
    			attr_dev(form, "action", "/login");
    			attr_dev(form, "method", "post");
    			attr_dev(form, "class", "card-content");
    			add_location(form, file$5, 24, 4, 655);
    			attr_dev(login_1, "class", "card");
    			add_location(login_1, file$5, 22, 0, 589);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, login_1, anchor);
    			append_dev(login_1, form);
    			append_dev(form, div1);
    			append_dev(div1, label0);
    			append_dev(div1, t1);
    			append_dev(div1, div0);
    			append_dev(div0, input0);
    			set_input_value(input0, /*username*/ ctx[0]);
    			append_dev(div0, t2);
    			append_dev(div0, span0);
    			append_dev(span0, i0);
    			append_dev(form, t3);
    			append_dev(form, div3);
    			append_dev(div3, label1);
    			append_dev(div3, t5);
    			append_dev(div3, div2);
    			append_dev(div2, input1);
    			set_input_value(input1, /*password*/ ctx[1]);
    			append_dev(div2, t6);
    			append_dev(div2, span1);
    			append_dev(span1, i1);
    			append_dev(form, t7);
    			append_dev(form, div4);
    			append_dev(div4, button);
    			append_dev(button, strong);
    			append_dev(form, t9);
    			append_dev(form, div5);
    			if (if_block) if_block.m(div5, null);
    			input0.focus();

    			if (!mounted) {
    				dispose = [
    					listen_dev(input0, "input", /*input0_input_handler*/ ctx[5]),
    					listen_dev(input1, "input", /*input1_input_handler*/ ctx[6]),
    					listen_dev(button, "click", prevent_default(/*login*/ ctx[3]), false, true, false),
    					listen_dev(form, "submit", prevent_default(/*login*/ ctx[3]), false, true, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*username*/ 1 && input0.value !== /*username*/ ctx[0]) {
    				set_input_value(input0, /*username*/ ctx[0]);
    			}

    			if (dirty & /*password*/ 2 && input1.value !== /*password*/ ctx[1]) {
    				set_input_value(input1, /*password*/ ctx[1]);
    			}

    			if (/*err*/ ctx[2]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);
    				} else {
    					if_block = create_if_block$4(ctx);
    					if_block.c();
    					if_block.m(div5, null);
    				}
    			} else if (if_block) {
    				if_block.d(1);
    				if_block = null;
    			}
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(login_1);
    			if (if_block) if_block.d();
    			mounted = false;
    			run_all(dispose);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$6.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$6($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots("Login", slots, []);
    	let { connectedStatus = "" } = $$props;
    	let username;
    	let password;
    	let err;

    	async function login() {
    		let response = await fetch("/login", {
    			method: "POST",
    			body: "username=" + username + "&password=" + password,
    			headers: {
    				"Content-Type": "application/x-www-form-urlencoded"
    			}
    		});

    		if (response.ok) {
    			$$invalidate(4, connectedStatus = true);
    		} else {
    			$$invalidate(4, connectedStatus = false);
    			$$invalidate(2, err = "Echec de la connexion. Veuillez vérifier votre identifiant et votre mot de passe. Si le problème persiste, merci de contacter l'administrateur.");
    		}
    	}

    	const writable_props = ["connectedStatus"];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== "$$") console.warn(`<Login> was created with unknown prop '${key}'`);
    	});

    	function input0_input_handler() {
    		username = this.value;
    		$$invalidate(0, username);
    	}

    	function input1_input_handler() {
    		password = this.value;
    		$$invalidate(1, password);
    	}

    	$$self.$$set = $$props => {
    		if ("connectedStatus" in $$props) $$invalidate(4, connectedStatus = $$props.connectedStatus);
    	};

    	$$self.$capture_state = () => ({
    		connectedStatus,
    		username,
    		password,
    		err,
    		login
    	});

    	$$self.$inject_state = $$props => {
    		if ("connectedStatus" in $$props) $$invalidate(4, connectedStatus = $$props.connectedStatus);
    		if ("username" in $$props) $$invalidate(0, username = $$props.username);
    		if ("password" in $$props) $$invalidate(1, password = $$props.password);
    		if ("err" in $$props) $$invalidate(2, err = $$props.err);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [
    		username,
    		password,
    		err,
    		login,
    		connectedStatus,
    		input0_input_handler,
    		input1_input_handler
    	];
    }

    class Login extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$6, create_fragment$6, safe_not_equal, { connectedStatus: 4 });

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Login",
    			options,
    			id: create_fragment$6.name
    		});
    	}

    	get connectedStatus() {
    		throw new Error("<Login>: Props cannot be read directly from the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}

    	set connectedStatus(value) {
    		throw new Error("<Login>: Props cannot be set directly on the component instance unless compiling with 'accessors: true' or '<svelte:options accessors/>'");
    	}
    }

    async function callApi(url) {
        try {
            let response = await fetch(url);
            if(response.ok) {
                console.log("debug");
                console.log(response);
                return response.json();
            }
            console.log("Erreur http " + response.status + " sur url " + url);
            return null;
        } catch (error) {
            console.log("Erreur réseau " + error);
            return null;
        }
    }

    async function callCheckConnexion() {
        let response = await fetch("check");
        return response.ok;
    }

    async function callBonjour(nom) {
        return await callApi("/hello?nom="+nom);
    }
    async function callMaj(nom) {
        return await callApi("/upper?nom="+nom);
    }
    async function callMin(nom) {
        return await callApi("/lower?nom="+nom);
    }
    async function callHistoric() {
        return await callApi("/historic");
    }

    var net = /*#__PURE__*/Object.freeze({
        __proto__: null,
        callApi: callApi,
        callCheckConnexion: callCheckConnexion,
        callBonjour: callBonjour,
        callMaj: callMaj,
        callMin: callMin,
        callHistoric: callHistoric
    });

    /* functions/Bonjour.svelte generated by Svelte v3.37.0 */
    const file$4 = "functions/Bonjour.svelte";

    // (19:4) {#if resultat}
    function create_if_block$3(ctx) {
    	let t;

    	const block = {
    		c: function create() {
    			t = text(/*resultat*/ ctx[1]);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, t, anchor);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*resultat*/ 2) set_data_dev(t, /*resultat*/ ctx[1]);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(t);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$3.name,
    		type: "if",
    		source: "(19:4) {#if resultat}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$5(ctx) {
    	let div;
    	let input;
    	let t0;
    	let button;
    	let t2;
    	let hr;
    	let t3;
    	let mounted;
    	let dispose;
    	let if_block = /*resultat*/ ctx[1] && create_if_block$3(ctx);

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			t0 = space();
    			button = element("button");
    			button.textContent = "Bonjour";
    			t2 = space();
    			hr = element("hr");
    			t3 = space();
    			if (if_block) if_block.c();
    			attr_dev(input, "type", "text");
    			attr_dev(input, "placeholder", "Entrer un prénom...");
    			attr_dev(input, "class", "input");
    			add_location(input, file$4, 12, 4, 232);
    			add_location(button, file$4, 14, 4, 331);
    			add_location(hr, file$4, 17, 4, 395);
    			add_location(div, file$4, 11, 0, 222);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			set_input_value(input, /*nom*/ ctx[0]);
    			append_dev(div, t0);
    			append_dev(div, button);
    			append_dev(div, t2);
    			append_dev(div, hr);
    			append_dev(div, t3);
    			if (if_block) if_block.m(div, null);

    			if (!mounted) {
    				dispose = [
    					listen_dev(input, "input", /*input_input_handler*/ ctx[3]),
    					listen_dev(button, "click", /*bonjour*/ ctx[2], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*nom*/ 1 && input.value !== /*nom*/ ctx[0]) {
    				set_input_value(input, /*nom*/ ctx[0]);
    			}

    			if (/*resultat*/ ctx[1]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);
    				} else {
    					if_block = create_if_block$3(ctx);
    					if_block.c();
    					if_block.m(div, null);
    				}
    			} else if (if_block) {
    				if_block.d(1);
    				if_block = null;
    			}
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			if (if_block) if_block.d();
    			mounted = false;
    			run_all(dispose);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$5.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$5($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots("Bonjour", slots, []);
    	let nom;
    	let resultat;

    	async function bonjour() {
    		let res = await callBonjour(nom);
    		if (res != null) $$invalidate(1, resultat = res);
    	}

    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== "$$") console.warn(`<Bonjour> was created with unknown prop '${key}'`);
    	});

    	function input_input_handler() {
    		nom = this.value;
    		$$invalidate(0, nom);
    	}

    	$$self.$capture_state = () => ({ net, nom, resultat, bonjour });

    	$$self.$inject_state = $$props => {
    		if ("nom" in $$props) $$invalidate(0, nom = $$props.nom);
    		if ("resultat" in $$props) $$invalidate(1, resultat = $$props.resultat);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [nom, resultat, bonjour, input_input_handler];
    }

    class Bonjour extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$5, create_fragment$5, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Bonjour",
    			options,
    			id: create_fragment$5.name
    		});
    	}
    }

    /* functions/Maj.svelte generated by Svelte v3.37.0 */
    const file$3 = "functions/Maj.svelte";

    // (19:4) {#if resultat}
    function create_if_block$2(ctx) {
    	let t;

    	const block = {
    		c: function create() {
    			t = text(/*resultat*/ ctx[1]);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, t, anchor);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*resultat*/ 2) set_data_dev(t, /*resultat*/ ctx[1]);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(t);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$2.name,
    		type: "if",
    		source: "(19:4) {#if resultat}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$4(ctx) {
    	let div;
    	let input;
    	let t0;
    	let button;
    	let t2;
    	let hr;
    	let t3;
    	let mounted;
    	let dispose;
    	let if_block = /*resultat*/ ctx[1] && create_if_block$2(ctx);

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			t0 = space();
    			button = element("button");
    			button.textContent = "Majuscule";
    			t2 = space();
    			hr = element("hr");
    			t3 = space();
    			if (if_block) if_block.c();
    			attr_dev(input, "type", "text");
    			attr_dev(input, "placeholder", "Entrer le mot ou la phrase à mettre en majuscule...");
    			attr_dev(input, "class", "input");
    			add_location(input, file$3, 12, 4, 224);
    			add_location(button, file$3, 14, 4, 355);
    			add_location(hr, file$3, 17, 4, 417);
    			add_location(div, file$3, 11, 0, 214);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			set_input_value(input, /*nom*/ ctx[0]);
    			append_dev(div, t0);
    			append_dev(div, button);
    			append_dev(div, t2);
    			append_dev(div, hr);
    			append_dev(div, t3);
    			if (if_block) if_block.m(div, null);

    			if (!mounted) {
    				dispose = [
    					listen_dev(input, "input", /*input_input_handler*/ ctx[3]),
    					listen_dev(button, "click", /*maj*/ ctx[2], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*nom*/ 1 && input.value !== /*nom*/ ctx[0]) {
    				set_input_value(input, /*nom*/ ctx[0]);
    			}

    			if (/*resultat*/ ctx[1]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);
    				} else {
    					if_block = create_if_block$2(ctx);
    					if_block.c();
    					if_block.m(div, null);
    				}
    			} else if (if_block) {
    				if_block.d(1);
    				if_block = null;
    			}
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			if (if_block) if_block.d();
    			mounted = false;
    			run_all(dispose);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$4.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$4($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots("Maj", slots, []);
    	let nom;
    	let resultat;

    	async function maj() {
    		let res = await callMaj(nom);
    		if (res != null) $$invalidate(1, resultat = res);
    	}

    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== "$$") console.warn(`<Maj> was created with unknown prop '${key}'`);
    	});

    	function input_input_handler() {
    		nom = this.value;
    		$$invalidate(0, nom);
    	}

    	$$self.$capture_state = () => ({ net, nom, resultat, maj });

    	$$self.$inject_state = $$props => {
    		if ("nom" in $$props) $$invalidate(0, nom = $$props.nom);
    		if ("resultat" in $$props) $$invalidate(1, resultat = $$props.resultat);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [nom, resultat, maj, input_input_handler];
    }

    class Maj extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$4, create_fragment$4, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Maj",
    			options,
    			id: create_fragment$4.name
    		});
    	}
    }

    /* functions/Min.svelte generated by Svelte v3.37.0 */
    const file$2 = "functions/Min.svelte";

    // (19:4) {#if resultat}
    function create_if_block$1(ctx) {
    	let t;

    	const block = {
    		c: function create() {
    			t = text(/*resultat*/ ctx[1]);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, t, anchor);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*resultat*/ 2) set_data_dev(t, /*resultat*/ ctx[1]);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(t);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block$1.name,
    		type: "if",
    		source: "(19:4) {#if resultat}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$3(ctx) {
    	let div;
    	let input;
    	let t0;
    	let button;
    	let t2;
    	let hr;
    	let t3;
    	let mounted;
    	let dispose;
    	let if_block = /*resultat*/ ctx[1] && create_if_block$1(ctx);

    	const block = {
    		c: function create() {
    			div = element("div");
    			input = element("input");
    			t0 = space();
    			button = element("button");
    			button.textContent = "Minuscule";
    			t2 = space();
    			hr = element("hr");
    			t3 = space();
    			if (if_block) if_block.c();
    			attr_dev(input, "type", "text");
    			attr_dev(input, "placeholder", "Entrer le mot ou la phrase à mettre en minuscule...");
    			attr_dev(input, "class", "input");
    			add_location(input, file$2, 12, 4, 224);
    			add_location(button, file$2, 14, 4, 355);
    			add_location(hr, file$2, 17, 4, 417);
    			add_location(div, file$2, 11, 0, 214);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, input);
    			set_input_value(input, /*nom*/ ctx[0]);
    			append_dev(div, t0);
    			append_dev(div, button);
    			append_dev(div, t2);
    			append_dev(div, hr);
    			append_dev(div, t3);
    			if (if_block) if_block.m(div, null);

    			if (!mounted) {
    				dispose = [
    					listen_dev(input, "input", /*input_input_handler*/ ctx[3]),
    					listen_dev(button, "click", /*min*/ ctx[2], false, false, false)
    				];

    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*nom*/ 1 && input.value !== /*nom*/ ctx[0]) {
    				set_input_value(input, /*nom*/ ctx[0]);
    			}

    			if (/*resultat*/ ctx[1]) {
    				if (if_block) {
    					if_block.p(ctx, dirty);
    				} else {
    					if_block = create_if_block$1(ctx);
    					if_block.c();
    					if_block.m(div, null);
    				}
    			} else if (if_block) {
    				if_block.d(1);
    				if_block = null;
    			}
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			if (if_block) if_block.d();
    			mounted = false;
    			run_all(dispose);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$3.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$3($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots("Min", slots, []);
    	let nom;
    	let resultat;

    	async function min() {
    		let res = await callMin(nom);
    		if (res != null) $$invalidate(1, resultat = res);
    	}

    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== "$$") console.warn(`<Min> was created with unknown prop '${key}'`);
    	});

    	function input_input_handler() {
    		nom = this.value;
    		$$invalidate(0, nom);
    	}

    	$$self.$capture_state = () => ({ net, nom, resultat, min });

    	$$self.$inject_state = $$props => {
    		if ("nom" in $$props) $$invalidate(0, nom = $$props.nom);
    		if ("resultat" in $$props) $$invalidate(1, resultat = $$props.resultat);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [nom, resultat, min, input_input_handler];
    }

    class Min extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$3, create_fragment$3, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Min",
    			options,
    			id: create_fragment$3.name
    		});
    	}
    }

    /* functions/Historic.svelte generated by Svelte v3.37.0 */
    const file$1 = "functions/Historic.svelte";

    function get_each_context$1(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[2] = list[i];
    	return child_ctx;
    }

    // (16:8) {#each historique as evenement}
    function create_each_block$1(ctx) {
    	let li;
    	let t_value = /*evenement*/ ctx[2] + "";
    	let t;

    	const block = {
    		c: function create() {
    			li = element("li");
    			t = text(t_value);
    			add_location(li, file$1, 16, 12, 336);
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, li, anchor);
    			append_dev(li, t);
    		},
    		p: function update(ctx, dirty) {
    			if (dirty & /*historique*/ 1 && t_value !== (t_value = /*evenement*/ ctx[2] + "")) set_data_dev(t, t_value);
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(li);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block$1.name,
    		type: "each",
    		source: "(16:8) {#each historique as evenement}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$2(ctx) {
    	let div;
    	let h1;
    	let t1;
    	let ul;
    	let each_value = /*historique*/ ctx[0];
    	validate_each_argument(each_value);
    	let each_blocks = [];

    	for (let i = 0; i < each_value.length; i += 1) {
    		each_blocks[i] = create_each_block$1(get_each_context$1(ctx, each_value, i));
    	}

    	const block = {
    		c: function create() {
    			div = element("div");
    			h1 = element("h1");
    			h1.textContent = "Historique";
    			t1 = space();
    			ul = element("ul");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			add_location(h1, file$1, 13, 4, 255);
    			add_location(ul, file$1, 14, 4, 279);
    			add_location(div, file$1, 12, 0, 245);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, div, anchor);
    			append_dev(div, h1);
    			append_dev(div, t1);
    			append_dev(div, ul);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(ul, null);
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*historique*/ 1) {
    				each_value = /*historique*/ ctx[0];
    				validate_each_argument(each_value);
    				let i;

    				for (i = 0; i < each_value.length; i += 1) {
    					const child_ctx = get_each_context$1(ctx, each_value, i);

    					if (each_blocks[i]) {
    						each_blocks[i].p(child_ctx, dirty);
    					} else {
    						each_blocks[i] = create_each_block$1(child_ctx);
    						each_blocks[i].c();
    						each_blocks[i].m(ul, null);
    					}
    				}

    				for (; i < each_blocks.length; i += 1) {
    					each_blocks[i].d(1);
    				}

    				each_blocks.length = each_value.length;
    			}
    		},
    		i: noop,
    		o: noop,
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(div);
    			destroy_each(each_blocks, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_fragment$2.name,
    		type: "component",
    		source: "",
    		ctx
    	});

    	return block;
    }

    function instance$2($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots("Historic", slots, []);
    	let historique = [];

    	async function reloadHistoric() {
    		let res = await callHistoric();
    		if (res != null) $$invalidate(0, historique = res);
    	}

    	reloadHistoric();
    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== "$$") console.warn(`<Historic> was created with unknown prop '${key}'`);
    	});

    	$$self.$capture_state = () => ({ net, historique, reloadHistoric });

    	$$self.$inject_state = $$props => {
    		if ("historique" in $$props) $$invalidate(0, historique = $$props.historique);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [historique];
    }

    class Historic extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$2, create_fragment$2, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Historic",
    			options,
    			id: create_fragment$2.name
    		});
    	}
    }

    /* Dashboard.svelte generated by Svelte v3.37.0 */
    const file = "Dashboard.svelte";

    function get_each_context(ctx, list, i) {
    	const child_ctx = ctx.slice();
    	child_ctx[6] = list[i];
    	return child_ctx;
    }

    // (38:3) {#each Menu as item (item.id)}
    function create_each_block(key_1, ctx) {
    	let a;
    	let span0;
    	let i;
    	let t0;
    	let span1;
    	let t1_value = /*item*/ ctx[6].label + "";
    	let t1;
    	let t2;
    	let mounted;
    	let dispose;

    	function click_handler_1() {
    		return /*click_handler_1*/ ctx[5](/*item*/ ctx[6]);
    	}

    	const block = {
    		key: key_1,
    		first: null,
    		c: function create() {
    			a = element("a");
    			span0 = element("span");
    			i = element("i");
    			t0 = space();
    			span1 = element("span");
    			t1 = text(t1_value);
    			t2 = space();
    			attr_dev(i, "class", "fas " + /*item*/ ctx[6].icon);
    			add_location(i, file, 39, 35, 1375);
    			attr_dev(span0, "class", "icon is-medium");
    			add_location(span0, file, 39, 5, 1345);
    			add_location(span1, file, 40, 5, 1420);
    			attr_dev(a, "href", "#" + /*item*/ ctx[6].id);
    			attr_dev(a, "class", "navbar-item");
    			add_location(a, file, 38, 4, 1263);
    			this.first = a;
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, a, anchor);
    			append_dev(a, span0);
    			append_dev(span0, i);
    			append_dev(a, t0);
    			append_dev(a, span1);
    			append_dev(span1, t1);
    			append_dev(a, t2);

    			if (!mounted) {
    				dispose = listen_dev(a, "click", click_handler_1, false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(new_ctx, dirty) {
    			ctx = new_ctx;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(a);
    			mounted = false;
    			dispose();
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_each_block.name,
    		type: "each",
    		source: "(38:3) {#each Menu as item (item.id)}",
    		ctx
    	});

    	return block;
    }

    function create_fragment$1(ctx) {
    	let nav;
    	let div2;
    	let div0;
    	let span0;
    	let i;
    	let t0;
    	let span1;
    	let t2;
    	let div1;
    	let span2;
    	let t3;
    	let span3;
    	let t4;
    	let span4;
    	let t5;
    	let div7;
    	let div3;
    	let each_blocks = [];
    	let each_1_lookup = new Map();
    	let t6;
    	let div6;
    	let div5;
    	let div4;
    	let a;
    	let t8;
    	let main;
    	let switch_instance;
    	let current;
    	let mounted;
    	let dispose;
    	let each_value = /*Menu*/ ctx[2];
    	validate_each_argument(each_value);
    	const get_key = ctx => /*item*/ ctx[6].id;
    	validate_each_keys(ctx, each_value, get_each_context, get_key);

    	for (let i = 0; i < each_value.length; i += 1) {
    		let child_ctx = get_each_context(ctx, each_value, i);
    		let key = get_key(child_ctx);
    		each_1_lookup.set(key, each_blocks[i] = create_each_block(key, child_ctx));
    	}

    	var switch_value = /*actualMenu*/ ctx[1].component;

    	function switch_props(ctx) {
    		return { $$inline: true };
    	}

    	if (switch_value) {
    		switch_instance = new switch_value(switch_props());
    	}

    	const block = {
    		c: function create() {
    			nav = element("nav");
    			div2 = element("div");
    			div0 = element("div");
    			span0 = element("span");
    			i = element("i");
    			t0 = space();
    			span1 = element("span");
    			span1.textContent = "Webtoolkit";
    			t2 = space();
    			div1 = element("div");
    			span2 = element("span");
    			t3 = space();
    			span3 = element("span");
    			t4 = space();
    			span4 = element("span");
    			t5 = space();
    			div7 = element("div");
    			div3 = element("div");

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].c();
    			}

    			t6 = space();
    			div6 = element("div");
    			div5 = element("div");
    			div4 = element("div");
    			a = element("a");
    			a.textContent = "Déconnexion";
    			t8 = space();
    			main = element("main");
    			if (switch_instance) create_component(switch_instance.$$.fragment);
    			attr_dev(i, "class", "fas fa-gem");
    			add_location(i, file, 25, 4, 899);
    			attr_dev(span0, "class", "icon is-medium");
    			add_location(span0, file, 24, 3, 865);
    			add_location(span1, file, 27, 3, 940);
    			attr_dev(div0, "class", "navbar-item is-size-5");
    			add_location(div0, file, 23, 2, 826);
    			add_location(span2, file, 30, 3, 1069);
    			add_location(span3, file, 31, 3, 1086);
    			add_location(span4, file, 32, 3, 1103);
    			attr_dev(div1, "class", "navbar-burger is-hoverable");
    			add_location(div1, file, 29, 2, 975);
    			attr_dev(div2, "class", "navbar-brand");
    			add_location(div2, file, 22, 1, 797);
    			attr_dev(div3, "class", "navbar-start");
    			add_location(div3, file, 36, 2, 1198);
    			attr_dev(a, "href", "/logout");
    			attr_dev(a, "class", "button");
    			add_location(a, file, 47, 5, 1562);
    			attr_dev(div4, "class", "buttons");
    			add_location(div4, file, 46, 4, 1535);
    			attr_dev(div5, "class", "navbar-item");
    			add_location(div5, file, 45, 3, 1505);
    			attr_dev(div6, "class", "navbar-end");
    			add_location(div6, file, 44, 2, 1477);
    			attr_dev(div7, "class", "navbar-menu");
    			toggle_class(div7, "is-active", /*showNavbarMenu*/ ctx[0]);
    			add_location(div7, file, 35, 1, 1135);
    			attr_dev(nav, "class", "navbar is-fixed-top is-wrapped has-shadow is-light");
    			add_location(nav, file, 21, 0, 731);
    			attr_dev(main, "class", "content is-wrapped");
    			add_location(main, file, 54, 0, 1657);
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			insert_dev(target, nav, anchor);
    			append_dev(nav, div2);
    			append_dev(div2, div0);
    			append_dev(div0, span0);
    			append_dev(span0, i);
    			append_dev(div0, t0);
    			append_dev(div0, span1);
    			append_dev(div2, t2);
    			append_dev(div2, div1);
    			append_dev(div1, span2);
    			append_dev(div1, t3);
    			append_dev(div1, span3);
    			append_dev(div1, t4);
    			append_dev(div1, span4);
    			append_dev(nav, t5);
    			append_dev(nav, div7);
    			append_dev(div7, div3);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].m(div3, null);
    			}

    			append_dev(div7, t6);
    			append_dev(div7, div6);
    			append_dev(div6, div5);
    			append_dev(div5, div4);
    			append_dev(div4, a);
    			insert_dev(target, t8, anchor);
    			insert_dev(target, main, anchor);

    			if (switch_instance) {
    				mount_component(switch_instance, main, null);
    			}

    			current = true;

    			if (!mounted) {
    				dispose = listen_dev(div1, "click", /*click_handler*/ ctx[4], false, false, false);
    				mounted = true;
    			}
    		},
    		p: function update(ctx, [dirty]) {
    			if (dirty & /*Menu, activate*/ 12) {
    				each_value = /*Menu*/ ctx[2];
    				validate_each_argument(each_value);
    				validate_each_keys(ctx, each_value, get_each_context, get_key);
    				each_blocks = update_keyed_each(each_blocks, dirty, get_key, 1, ctx, each_value, each_1_lookup, div3, destroy_block, create_each_block, null, get_each_context);
    			}

    			if (dirty & /*showNavbarMenu*/ 1) {
    				toggle_class(div7, "is-active", /*showNavbarMenu*/ ctx[0]);
    			}

    			if (switch_value !== (switch_value = /*actualMenu*/ ctx[1].component)) {
    				if (switch_instance) {
    					group_outros();
    					const old_component = switch_instance;

    					transition_out(old_component.$$.fragment, 1, 0, () => {
    						destroy_component(old_component, 1);
    					});

    					check_outros();
    				}

    				if (switch_value) {
    					switch_instance = new switch_value(switch_props());
    					create_component(switch_instance.$$.fragment);
    					transition_in(switch_instance.$$.fragment, 1);
    					mount_component(switch_instance, main, null);
    				} else {
    					switch_instance = null;
    				}
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			if (switch_instance) transition_in(switch_instance.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			if (switch_instance) transition_out(switch_instance.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if (detaching) detach_dev(nav);

    			for (let i = 0; i < each_blocks.length; i += 1) {
    				each_blocks[i].d();
    			}

    			if (detaching) detach_dev(t8);
    			if (detaching) detach_dev(main);
    			if (switch_instance) destroy_component(switch_instance);
    			mounted = false;
    			dispose();
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

    function instance$1($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots("Dashboard", slots, []);

    	let Menu = [
    		{
    			id: "bonjour",
    			component: Bonjour,
    			icon: "fa-home",
    			label: "Bonjour"
    		},
    		{
    			id: "maj",
    			component: Maj,
    			icon: "fa-chart-bar",
    			label: "Majuscule"
    		},
    		{
    			id: "min",
    			component: Min,
    			icon: "fa-book",
    			label: "Minuscule"
    		},
    		{
    			id: "historic",
    			component: Historic,
    			icon: "fa-shopping-cart",
    			label: "Historique"
    		}
    	];

    	let showNavbarMenu = false;
    	let actualMenu = Menu[0];

    	function activate(menuID) {
    		$$invalidate(1, actualMenu = Menu.find(elem => elem.id === menuID));
    	}

    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== "$$") console.warn(`<Dashboard> was created with unknown prop '${key}'`);
    	});

    	const click_handler = () => $$invalidate(0, showNavbarMenu = !showNavbarMenu);
    	const click_handler_1 = item => activate(item.id);

    	$$self.$capture_state = () => ({
    		Bonjour,
    		Maj,
    		Min,
    		Historic,
    		Menu,
    		showNavbarMenu,
    		actualMenu,
    		activate
    	});

    	$$self.$inject_state = $$props => {
    		if ("Menu" in $$props) $$invalidate(2, Menu = $$props.Menu);
    		if ("showNavbarMenu" in $$props) $$invalidate(0, showNavbarMenu = $$props.showNavbarMenu);
    		if ("actualMenu" in $$props) $$invalidate(1, actualMenu = $$props.actualMenu);
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [showNavbarMenu, actualMenu, Menu, activate, click_handler, click_handler_1];
    }

    class Dashboard extends SvelteComponentDev {
    	constructor(options) {
    		super(options);
    		init(this, options, instance$1, create_fragment$1, safe_not_equal, {});

    		dispatch_dev("SvelteRegisterComponent", {
    			component: this,
    			tagName: "Dashboard",
    			options,
    			id: create_fragment$1.name
    		});
    	}
    }

    class Cookie {
        get(cname) {
            var name = cname + "=";
            var decodedCookie = decodeURIComponent(document.cookie);
            var ca = decodedCookie.split(';');
            for(var i = 0; i <ca.length; i++) {
                var c = ca[i];
                while (c.charAt(0) == ' ') {
                    c = c.substring(1);
                }
                if (c.indexOf(name) == 0) {
                    return c.substring(name.length, c.length);
                }
            }
            return "";
        }
    }

    /* App.svelte generated by Svelte v3.37.0 */

    // (3:0) {:else}
    function create_else_block(ctx) {
    	let dashboard;
    	let current;
    	dashboard = new Dashboard({ $$inline: true });

    	const block = {
    		c: function create() {
    			create_component(dashboard.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(dashboard, target, anchor);
    			current = true;
    		},
    		p: noop,
    		i: function intro(local) {
    			if (current) return;
    			transition_in(dashboard.$$.fragment, local);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(dashboard.$$.fragment, local);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			destroy_component(dashboard, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_else_block.name,
    		type: "else",
    		source: "(3:0) {:else}",
    		ctx
    	});

    	return block;
    }

    // (1:0) {#if !connectedStatus}
    function create_if_block(ctx) {
    	let login;
    	let updating_connectedStatus;
    	let current;

    	function login_connectedStatus_binding(value) {
    		/*login_connectedStatus_binding*/ ctx[1](value);
    	}

    	let login_props = {};

    	if (/*connectedStatus*/ ctx[0] !== void 0) {
    		login_props.connectedStatus = /*connectedStatus*/ ctx[0];
    	}

    	login = new Login({ props: login_props, $$inline: true });
    	binding_callbacks.push(() => bind(login, "connectedStatus", login_connectedStatus_binding));

    	const block = {
    		c: function create() {
    			create_component(login.$$.fragment);
    		},
    		m: function mount(target, anchor) {
    			mount_component(login, target, anchor);
    			current = true;
    		},
    		p: function update(ctx, dirty) {
    			const login_changes = {};

    			if (!updating_connectedStatus && dirty & /*connectedStatus*/ 1) {
    				updating_connectedStatus = true;
    				login_changes.connectedStatus = /*connectedStatus*/ ctx[0];
    				add_flush_callback(() => updating_connectedStatus = false);
    			}

    			login.$set(login_changes);
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
    			destroy_component(login, detaching);
    		}
    	};

    	dispatch_dev("SvelteRegisterBlock", {
    		block,
    		id: create_if_block.name,
    		type: "if",
    		source: "(1:0) {#if !connectedStatus}",
    		ctx
    	});

    	return block;
    }

    function create_fragment(ctx) {
    	let current_block_type_index;
    	let if_block;
    	let if_block_anchor;
    	let current;
    	const if_block_creators = [create_if_block, create_else_block];
    	const if_blocks = [];

    	function select_block_type(ctx, dirty) {
    		if (!/*connectedStatus*/ ctx[0]) return 0;
    		return 1;
    	}

    	current_block_type_index = select_block_type(ctx);
    	if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);

    	const block = {
    		c: function create() {
    			if_block.c();
    			if_block_anchor = empty();
    		},
    		l: function claim(nodes) {
    			throw new Error("options.hydrate only works if the component was compiled with the `hydratable: true` option");
    		},
    		m: function mount(target, anchor) {
    			if_blocks[current_block_type_index].m(target, anchor);
    			insert_dev(target, if_block_anchor, anchor);
    			current = true;
    		},
    		p: function update(ctx, [dirty]) {
    			let previous_block_index = current_block_type_index;
    			current_block_type_index = select_block_type(ctx);

    			if (current_block_type_index === previous_block_index) {
    				if_blocks[current_block_type_index].p(ctx, dirty);
    			} else {
    				group_outros();

    				transition_out(if_blocks[previous_block_index], 1, 1, () => {
    					if_blocks[previous_block_index] = null;
    				});

    				check_outros();
    				if_block = if_blocks[current_block_type_index];

    				if (!if_block) {
    					if_block = if_blocks[current_block_type_index] = if_block_creators[current_block_type_index](ctx);
    					if_block.c();
    				} else {
    					if_block.p(ctx, dirty);
    				}

    				transition_in(if_block, 1);
    				if_block.m(if_block_anchor.parentNode, if_block_anchor);
    			}
    		},
    		i: function intro(local) {
    			if (current) return;
    			transition_in(if_block);
    			current = true;
    		},
    		o: function outro(local) {
    			transition_out(if_block);
    			current = false;
    		},
    		d: function destroy(detaching) {
    			if_blocks[current_block_type_index].d(detaching);
    			if (detaching) detach_dev(if_block_anchor);
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

    function instance($$self, $$props, $$invalidate) {
    	let { $$slots: slots = {}, $$scope } = $$props;
    	validate_slots("App", slots, []);
    	let connectedStatus = false;
    	let idCookieValue = new Cookie().get("jeton");

    	if (idCookieValue === "") {
    		// pas de jeton => l'utilisateur n'est pas connecté
    		connectedStatus = false;
    	} else {
    		// si on a un jeton, on vérifie auprès du serveur
    		callCheckConnexion().then(response => {
    			$$invalidate(0, connectedStatus = response);
    		});
    	}

    	const writable_props = [];

    	Object.keys($$props).forEach(key => {
    		if (!~writable_props.indexOf(key) && key.slice(0, 2) !== "$$") console.warn(`<App> was created with unknown prop '${key}'`);
    	});

    	function login_connectedStatus_binding(value) {
    		connectedStatus = value;
    		$$invalidate(0, connectedStatus);
    	}

    	$$self.$capture_state = () => ({
    		Login,
    		Dashboard,
    		Cookie,
    		net,
    		connectedStatus,
    		idCookieValue
    	});

    	$$self.$inject_state = $$props => {
    		if ("connectedStatus" in $$props) $$invalidate(0, connectedStatus = $$props.connectedStatus);
    		if ("idCookieValue" in $$props) idCookieValue = $$props.idCookieValue;
    	};

    	if ($$props && "$$inject" in $$props) {
    		$$self.$inject_state($$props.$$inject);
    	}

    	return [connectedStatus, login_connectedStatus_binding];
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
