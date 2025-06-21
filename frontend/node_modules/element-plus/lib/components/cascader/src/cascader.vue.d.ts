import { nextTick } from 'vue';
import type { ComputedRef, Ref } from 'vue';
import type { CascaderNode, CascaderValue } from 'element-plus/es/components/cascader-panel';
declare function __VLS_template(): {
    empty?(_: {}): any;
    empty?(_: {}): any;
};
declare const __VLS_component: import("vue").DefineComponent<{
    emptyValues: ArrayConstructor;
    valueOnClear: import("element-plus/es/utils").EpPropFinalized<readonly [StringConstructor, NumberConstructor, BooleanConstructor, FunctionConstructor], unknown, unknown, undefined, boolean>;
    size: {
        readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<StringConstructor, "" | "small" | "default" | "large", never>>;
        readonly required: false;
        readonly validator: ((val: unknown) => boolean) | undefined;
        __epPropKey: true;
    };
    placeholder: StringConstructor;
    disabled: BooleanConstructor;
    clearable: BooleanConstructor;
    filterable: BooleanConstructor;
    filterMethod: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => (node: CascaderNode, keyword: string) => boolean) | (() => (node: CascaderNode, keyword: string) => boolean) | {
        (): (node: CascaderNode, keyword: string) => boolean;
        new (): any;
        readonly prototype: any;
    } | ((new (...args: any[]) => (node: CascaderNode, keyword: string) => boolean) | (() => (node: CascaderNode, keyword: string) => boolean) | {
        (): (node: CascaderNode, keyword: string) => boolean;
        new (): any;
        readonly prototype: any;
    })[], unknown, unknown, (node: CascaderNode, keyword: string) => boolean, boolean>;
    separator: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
    showAllLevels: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
    collapseTags: BooleanConstructor;
    maxCollapseTags: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
    collapseTagsTooltip: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
    debounce: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
    beforeFilter: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => (value: string) => boolean | Promise<any>) | (() => (value: string) => boolean | Promise<any>) | {
        (): (value: string) => boolean | Promise<any>;
        new (): any;
        readonly prototype: any;
    } | ((new (...args: any[]) => (value: string) => boolean | Promise<any>) | (() => (value: string) => boolean | Promise<any>) | {
        (): (value: string) => boolean | Promise<any>;
        new (): any;
        readonly prototype: any;
    })[], unknown, unknown, () => true, boolean>;
    placement: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => "top" | "bottom" | "left" | "right" | "auto" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => import("element-plus/es/components/popper").Placement) | ((new (...args: any[]) => "top" | "bottom" | "left" | "right" | "auto" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => import("element-plus/es/components/popper").Placement))[], import("element-plus/es/components/popper").Placement, unknown, string, boolean>;
    fallbackPlacements: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/popper").Placement[]) | (() => import("element-plus/es/components/popper").Placement[]) | ((new (...args: any[]) => import("element-plus/es/components/popper").Placement[]) | (() => import("element-plus/es/components/popper").Placement[]))[], unknown, unknown, string[], boolean>;
    popperClass: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
    teleported: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
    tagType: {
        default: string;
        type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<StringConstructor, "success" | "warning" | "info" | "primary" | "danger", unknown>>;
        required: false;
        validator: ((val: unknown) => boolean) | undefined;
        __epPropKey: true;
    };
    tagEffect: {
        default: string;
        type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<StringConstructor, "dark" | "light" | "plain", unknown>>;
        required: false;
        validator: ((val: unknown) => boolean) | undefined;
        __epPropKey: true;
    };
    validateEvent: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
    persistent: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
    modelValue: {
        readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<(new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue) | ((new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue))[], unknown, unknown>>;
        readonly required: false;
        readonly validator: ((val: unknown) => boolean) | undefined;
        __epPropKey: true;
    };
    options: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderOption[], boolean>;
    props: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderProps, boolean>;
}, {
    /**
     * @description get an array of currently selected node,(leafOnly) whether only return the leaf checked nodes, default is `false`
     */
    getCheckedNodes: (leafOnly: boolean) => CascaderNode[] | undefined;
    /**
     * @description cascader panel ref
     */
    cascaderPanelRef: Ref<({
        $: import("vue").ComponentInternalInstance;
        $data: {};
        $props: Partial<{
            border: boolean;
            props: import("element-plus/es/components/cascader-panel").CascaderProps;
            options: import("element-plus/es/components/cascader-panel").CascaderOption[];
        }> & Omit<Readonly<import("vue").ExtractPropTypes<{
            border: {
                type: BooleanConstructor;
                default: boolean;
            };
            renderLabel: import("vue").PropType<import("../../cascader-panel/src/node").RenderLabel>;
            modelValue: {
                readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<(new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue) | ((new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue))[], unknown, unknown>>;
                readonly required: false;
                readonly validator: ((val: unknown) => boolean) | undefined;
                __epPropKey: true;
            };
            options: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderOption[], boolean>;
            props: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderProps, boolean>;
        }>> & {
            "onUpdate:modelValue"?: ((...args: any[]) => any) | undefined;
            onChange?: ((...args: any[]) => any) | undefined;
            onClose?: ((...args: any[]) => any) | undefined;
            "onExpand-change"?: ((...args: any[]) => any) | undefined;
        } & import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, "border" | "props" | "options">;
        $attrs: {
            [x: string]: unknown;
        };
        $refs: {
            [x: string]: unknown;
        };
        $slots: import("vue").Slots;
        $root: import("vue").ComponentPublicInstance | null;
        $parent: import("vue").ComponentPublicInstance | null;
        $emit: (event: "update:modelValue" | "change" | "close" | "expand-change", ...args: any[]) => void;
        $el: any;
        $options: import("vue").ComponentOptionsBase<Readonly<import("vue").ExtractPropTypes<{
            border: {
                type: BooleanConstructor;
                default: boolean;
            };
            renderLabel: import("vue").PropType<import("../../cascader-panel/src/node").RenderLabel>;
            modelValue: {
                readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<(new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue) | ((new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue))[], unknown, unknown>>;
                readonly required: false;
                readonly validator: ((val: unknown) => boolean) | undefined;
                __epPropKey: true;
            };
            options: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderOption[], boolean>;
            props: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderProps, boolean>;
        }>> & {
            "onUpdate:modelValue"?: ((...args: any[]) => any) | undefined;
            onChange?: ((...args: any[]) => any) | undefined;
            onClose?: ((...args: any[]) => any) | undefined;
            "onExpand-change"?: ((...args: any[]) => any) | undefined;
        }, {
            ns: {
                namespace: ComputedRef<string>;
                b: (blockSuffix?: string) => string;
                e: (element?: string) => string;
                m: (modifier?: string) => string;
                be: (blockSuffix?: string, element?: string) => string;
                em: (element?: string, modifier?: string) => string;
                bm: (blockSuffix?: string, modifier?: string) => string;
                bem: (blockSuffix?: string, element?: string, modifier?: string) => string;
                is: {
                    (name: string, state: boolean | undefined): string;
                    (name: string): string;
                };
                cssVar: (object: Record<string, string>) => Record<string, string>;
                cssVarName: (name: string) => string;
                cssVarBlock: (object: Record<string, string>) => Record<string, string>;
                cssVarBlockName: (name: string) => string;
            };
            menuList: Ref<any[]>;
            menus: Ref<{
                readonly uid: number;
                readonly level: number;
                readonly value: import("../../cascader-panel/src/node").CascaderNodeValue;
                readonly label: string;
                readonly pathNodes: any[];
                readonly pathValues: import("../../cascader-panel/src/node").CascaderNodeValue[];
                readonly pathLabels: string[];
                childrenData: {
                    [x: string]: unknown;
                    label?: string | undefined;
                    value?: import("../../cascader-panel/src/node").CascaderNodeValue | undefined;
                    children?: any[] | undefined;
                    disabled?: boolean | undefined;
                    leaf?: boolean | undefined;
                }[] | undefined;
                children: any[];
                text: string;
                loaded: boolean;
                checked: boolean;
                indeterminate: boolean;
                loading: boolean;
                readonly data: {
                    [x: string]: unknown;
                    label?: string | undefined;
                    value?: import("../../cascader-panel/src/node").CascaderNodeValue | undefined;
                    children?: any[] | undefined;
                    disabled?: boolean | undefined;
                    leaf?: boolean | undefined;
                } | null;
                readonly config: {
                    expandTrigger: import("element-plus/es/components/cascader-panel").ExpandTrigger;
                    multiple: boolean;
                    checkStrictly: boolean;
                    emitPath: boolean;
                    lazy: boolean;
                    lazyLoad: import("../../cascader-panel/src/node").LazyLoad;
                    value: string;
                    label: string;
                    children: string;
                    disabled: string | import("../../cascader-panel/src/node").isDisabled;
                    leaf: string | import("../../cascader-panel/src/node").isLeaf;
                    hoverThreshold: number;
                };
                readonly parent?: any | undefined;
                readonly root: boolean;
                readonly isDisabled: boolean;
                readonly isLeaf: boolean;
                readonly valueByOption: import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodeValue[];
                appendChild: (childData: import("element-plus/es/components/cascader-panel").CascaderOption) => CascaderNode;
                calcText: (allLevels: boolean, separator: string) => string;
                broadcast: (event: string, ...args: unknown[]) => void;
                emit: (event: string, ...args: unknown[]) => void;
                onParentCheck: (checked: boolean) => void;
                onChildCheck: () => void;
                setCheckState: (checked: boolean) => void;
                doCheck: (checked: boolean) => void;
            }[][]>;
            checkedNodes: Ref<{
                readonly uid: number;
                readonly level: number;
                readonly value: import("../../cascader-panel/src/node").CascaderNodeValue;
                readonly label: string;
                readonly pathNodes: any[];
                readonly pathValues: import("../../cascader-panel/src/node").CascaderNodeValue[];
                readonly pathLabels: string[];
                childrenData: {
                    [x: string]: unknown;
                    label?: string | undefined;
                    value?: import("../../cascader-panel/src/node").CascaderNodeValue | undefined;
                    children?: any[] | undefined;
                    disabled?: boolean | undefined;
                    leaf?: boolean | undefined;
                }[] | undefined;
                children: any[];
                text: string;
                loaded: boolean;
                checked: boolean;
                indeterminate: boolean;
                loading: boolean;
                readonly data: {
                    [x: string]: unknown;
                    label?: string | undefined;
                    value?: import("../../cascader-panel/src/node").CascaderNodeValue | undefined;
                    children?: any[] | undefined;
                    disabled?: boolean | undefined;
                    leaf?: boolean | undefined;
                } | null;
                readonly config: {
                    expandTrigger: import("element-plus/es/components/cascader-panel").ExpandTrigger;
                    multiple: boolean;
                    checkStrictly: boolean;
                    emitPath: boolean;
                    lazy: boolean;
                    lazyLoad: import("../../cascader-panel/src/node").LazyLoad;
                    value: string;
                    label: string;
                    children: string;
                    disabled: string | import("../../cascader-panel/src/node").isDisabled;
                    leaf: string | import("../../cascader-panel/src/node").isLeaf;
                    hoverThreshold: number;
                };
                readonly parent?: any | undefined;
                readonly root: boolean;
                readonly isDisabled: boolean;
                readonly isLeaf: boolean;
                readonly valueByOption: import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodeValue[];
                appendChild: (childData: import("element-plus/es/components/cascader-panel").CascaderOption) => CascaderNode;
                calcText: (allLevels: boolean, separator: string) => string;
                broadcast: (event: string, ...args: unknown[]) => void;
                emit: (event: string, ...args: unknown[]) => void;
                onParentCheck: (checked: boolean) => void;
                onChildCheck: () => void;
                setCheckState: (checked: boolean) => void;
                doCheck: (checked: boolean) => void;
            }[]>;
            handleKeyDown: (e: KeyboardEvent) => void;
            handleCheckChange: (node: CascaderNode, checked: boolean, emitClose?: boolean) => void;
            getFlattedNodes: (leafOnly: boolean) => CascaderNode[] | undefined;
            getCheckedNodes: (leafOnly: boolean) => CascaderNode[] | undefined;
            clearCheckedNodes: () => void;
            calculateCheckedValue: () => void;
            scrollToExpandingNode: () => void;
        }, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, ("update:modelValue" | "change" | "close" | "expand-change")[], string, {
            border: boolean;
            props: import("element-plus/es/components/cascader-panel").CascaderProps;
            options: import("element-plus/es/components/cascader-panel").CascaderOption[];
        }> & {
            beforeCreate?: (() => void) | (() => void)[];
            created?: (() => void) | (() => void)[];
            beforeMount?: (() => void) | (() => void)[];
            mounted?: (() => void) | (() => void)[];
            beforeUpdate?: (() => void) | (() => void)[];
            updated?: (() => void) | (() => void)[];
            activated?: (() => void) | (() => void)[];
            deactivated?: (() => void) | (() => void)[];
            beforeDestroy?: (() => void) | (() => void)[];
            beforeUnmount?: (() => void) | (() => void)[];
            destroyed?: (() => void) | (() => void)[];
            unmounted?: (() => void) | (() => void)[];
            renderTracked?: ((e: import("vue").DebuggerEvent) => void) | ((e: import("vue").DebuggerEvent) => void)[];
            renderTriggered?: ((e: import("vue").DebuggerEvent) => void) | ((e: import("vue").DebuggerEvent) => void)[];
            errorCaptured?: ((err: unknown, instance: import("vue").ComponentPublicInstance | null, info: string) => boolean | void) | ((err: unknown, instance: import("vue").ComponentPublicInstance | null, info: string) => boolean | void)[];
        };
        $forceUpdate: () => void;
        $nextTick: typeof nextTick;
        $watch(source: string | Function, cb: Function, options?: import("vue").WatchOptions): import("vue").WatchStopHandle;
    } & Readonly<import("vue").ExtractPropTypes<{
        border: {
            type: BooleanConstructor;
            default: boolean;
        };
        renderLabel: import("vue").PropType<import("../../cascader-panel/src/node").RenderLabel>;
        modelValue: {
            readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<(new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue) | ((new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue))[], unknown, unknown>>;
            readonly required: false;
            readonly validator: ((val: unknown) => boolean) | undefined;
            __epPropKey: true;
        };
        options: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderOption[], boolean>;
        props: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderProps, boolean>;
    }>> & {
        "onUpdate:modelValue"?: ((...args: any[]) => any) | undefined;
        onChange?: ((...args: any[]) => any) | undefined;
        onClose?: ((...args: any[]) => any) | undefined;
        "onExpand-change"?: ((...args: any[]) => any) | undefined;
    } & import("vue").ShallowUnwrapRef<{
        ns: {
            namespace: ComputedRef<string>;
            b: (blockSuffix?: string) => string;
            e: (element?: string) => string;
            m: (modifier?: string) => string;
            be: (blockSuffix?: string, element?: string) => string;
            em: (element?: string, modifier?: string) => string;
            bm: (blockSuffix?: string, modifier?: string) => string;
            bem: (blockSuffix?: string, element?: string, modifier?: string) => string;
            is: {
                (name: string, state: boolean | undefined): string;
                (name: string): string;
            };
            cssVar: (object: Record<string, string>) => Record<string, string>;
            cssVarName: (name: string) => string;
            cssVarBlock: (object: Record<string, string>) => Record<string, string>;
            cssVarBlockName: (name: string) => string;
        };
        menuList: Ref<any[]>;
        menus: Ref<{
            readonly uid: number;
            readonly level: number;
            readonly value: import("../../cascader-panel/src/node").CascaderNodeValue;
            readonly label: string;
            readonly pathNodes: any[];
            readonly pathValues: import("../../cascader-panel/src/node").CascaderNodeValue[];
            readonly pathLabels: string[];
            childrenData: {
                [x: string]: unknown;
                label?: string | undefined;
                value?: import("../../cascader-panel/src/node").CascaderNodeValue | undefined;
                children?: any[] | undefined;
                disabled?: boolean | undefined;
                leaf?: boolean | undefined;
            }[] | undefined;
            children: any[];
            text: string;
            loaded: boolean;
            checked: boolean;
            indeterminate: boolean;
            loading: boolean;
            readonly data: {
                [x: string]: unknown;
                label?: string | undefined;
                value?: import("../../cascader-panel/src/node").CascaderNodeValue | undefined;
                children?: any[] | undefined;
                disabled?: boolean | undefined;
                leaf?: boolean | undefined;
            } | null;
            readonly config: {
                expandTrigger: import("element-plus/es/components/cascader-panel").ExpandTrigger;
                multiple: boolean;
                checkStrictly: boolean;
                emitPath: boolean;
                lazy: boolean;
                lazyLoad: import("../../cascader-panel/src/node").LazyLoad;
                value: string;
                label: string;
                children: string;
                disabled: string | import("../../cascader-panel/src/node").isDisabled;
                leaf: string | import("../../cascader-panel/src/node").isLeaf;
                hoverThreshold: number;
            };
            readonly parent?: any | undefined;
            readonly root: boolean;
            readonly isDisabled: boolean;
            readonly isLeaf: boolean;
            readonly valueByOption: import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodeValue[];
            appendChild: (childData: import("element-plus/es/components/cascader-panel").CascaderOption) => CascaderNode;
            calcText: (allLevels: boolean, separator: string) => string;
            broadcast: (event: string, ...args: unknown[]) => void;
            emit: (event: string, ...args: unknown[]) => void;
            onParentCheck: (checked: boolean) => void;
            onChildCheck: () => void;
            setCheckState: (checked: boolean) => void;
            doCheck: (checked: boolean) => void;
        }[][]>;
        checkedNodes: Ref<{
            readonly uid: number;
            readonly level: number;
            readonly value: import("../../cascader-panel/src/node").CascaderNodeValue;
            readonly label: string;
            readonly pathNodes: any[];
            readonly pathValues: import("../../cascader-panel/src/node").CascaderNodeValue[];
            readonly pathLabels: string[];
            childrenData: {
                [x: string]: unknown;
                label?: string | undefined;
                value?: import("../../cascader-panel/src/node").CascaderNodeValue | undefined;
                children?: any[] | undefined;
                disabled?: boolean | undefined;
                leaf?: boolean | undefined;
            }[] | undefined;
            children: any[];
            text: string;
            loaded: boolean;
            checked: boolean;
            indeterminate: boolean;
            loading: boolean;
            readonly data: {
                [x: string]: unknown;
                label?: string | undefined;
                value?: import("../../cascader-panel/src/node").CascaderNodeValue | undefined;
                children?: any[] | undefined;
                disabled?: boolean | undefined;
                leaf?: boolean | undefined;
            } | null;
            readonly config: {
                expandTrigger: import("element-plus/es/components/cascader-panel").ExpandTrigger;
                multiple: boolean;
                checkStrictly: boolean;
                emitPath: boolean;
                lazy: boolean;
                lazyLoad: import("../../cascader-panel/src/node").LazyLoad;
                value: string;
                label: string;
                children: string;
                disabled: string | import("../../cascader-panel/src/node").isDisabled;
                leaf: string | import("../../cascader-panel/src/node").isLeaf;
                hoverThreshold: number;
            };
            readonly parent?: any | undefined;
            readonly root: boolean;
            readonly isDisabled: boolean;
            readonly isLeaf: boolean;
            readonly valueByOption: import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodeValue[];
            appendChild: (childData: import("element-plus/es/components/cascader-panel").CascaderOption) => CascaderNode;
            calcText: (allLevels: boolean, separator: string) => string;
            broadcast: (event: string, ...args: unknown[]) => void;
            emit: (event: string, ...args: unknown[]) => void;
            onParentCheck: (checked: boolean) => void;
            onChildCheck: () => void;
            setCheckState: (checked: boolean) => void;
            doCheck: (checked: boolean) => void;
        }[]>;
        handleKeyDown: (e: KeyboardEvent) => void;
        handleCheckChange: (node: CascaderNode, checked: boolean, emitClose?: boolean) => void;
        getFlattedNodes: (leafOnly: boolean) => CascaderNode[] | undefined;
        getCheckedNodes: (leafOnly: boolean) => CascaderNode[] | undefined;
        clearCheckedNodes: () => void;
        calculateCheckedValue: () => void;
        scrollToExpandingNode: () => void;
    }> & {} & import("vue").ComponentCustomProperties) | null>;
    /**
     * @description toggle the visible of popper
     */
    togglePopperVisible: (visible?: boolean) => void;
    /**
     * @description cascader content ref
     */
    contentRef: ComputedRef<HTMLElement | undefined>;
    /**
     * @description selected content text
     */
    presentText: ComputedRef<string>;
}, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
    clear: () => void;
    "update:modelValue": (_: CascaderValue) => void;
    change: (_: CascaderValue) => void;
    blur: (evt: FocusEvent) => void;
    focus: (evt: FocusEvent) => void;
    visibleChange: (val: boolean) => void;
    expandChange: (val: CascaderValue) => void;
    removeTag: (val: import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue) => void;
}, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    emptyValues: ArrayConstructor;
    valueOnClear: import("element-plus/es/utils").EpPropFinalized<readonly [StringConstructor, NumberConstructor, BooleanConstructor, FunctionConstructor], unknown, unknown, undefined, boolean>;
    size: {
        readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<StringConstructor, "" | "small" | "default" | "large", never>>;
        readonly required: false;
        readonly validator: ((val: unknown) => boolean) | undefined;
        __epPropKey: true;
    };
    placeholder: StringConstructor;
    disabled: BooleanConstructor;
    clearable: BooleanConstructor;
    filterable: BooleanConstructor;
    filterMethod: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => (node: CascaderNode, keyword: string) => boolean) | (() => (node: CascaderNode, keyword: string) => boolean) | {
        (): (node: CascaderNode, keyword: string) => boolean;
        new (): any;
        readonly prototype: any;
    } | ((new (...args: any[]) => (node: CascaderNode, keyword: string) => boolean) | (() => (node: CascaderNode, keyword: string) => boolean) | {
        (): (node: CascaderNode, keyword: string) => boolean;
        new (): any;
        readonly prototype: any;
    })[], unknown, unknown, (node: CascaderNode, keyword: string) => boolean, boolean>;
    separator: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
    showAllLevels: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
    collapseTags: BooleanConstructor;
    maxCollapseTags: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
    collapseTagsTooltip: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
    debounce: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, number, boolean>;
    beforeFilter: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => (value: string) => boolean | Promise<any>) | (() => (value: string) => boolean | Promise<any>) | {
        (): (value: string) => boolean | Promise<any>;
        new (): any;
        readonly prototype: any;
    } | ((new (...args: any[]) => (value: string) => boolean | Promise<any>) | (() => (value: string) => boolean | Promise<any>) | {
        (): (value: string) => boolean | Promise<any>;
        new (): any;
        readonly prototype: any;
    })[], unknown, unknown, () => true, boolean>;
    placement: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => "top" | "bottom" | "left" | "right" | "auto" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => import("element-plus/es/components/popper").Placement) | ((new (...args: any[]) => "top" | "bottom" | "left" | "right" | "auto" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => import("element-plus/es/components/popper").Placement))[], import("element-plus/es/components/popper").Placement, unknown, string, boolean>;
    fallbackPlacements: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/popper").Placement[]) | (() => import("element-plus/es/components/popper").Placement[]) | ((new (...args: any[]) => import("element-plus/es/components/popper").Placement[]) | (() => import("element-plus/es/components/popper").Placement[]))[], unknown, unknown, string[], boolean>;
    popperClass: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, string, boolean>;
    teleported: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
    tagType: {
        default: string;
        type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<StringConstructor, "success" | "warning" | "info" | "primary" | "danger", unknown>>;
        required: false;
        validator: ((val: unknown) => boolean) | undefined;
        __epPropKey: true;
    };
    tagEffect: {
        default: string;
        type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<StringConstructor, "dark" | "light" | "plain", unknown>>;
        required: false;
        validator: ((val: unknown) => boolean) | undefined;
        __epPropKey: true;
    };
    validateEvent: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
    persistent: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, boolean, boolean>;
    modelValue: {
        readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<(new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue) | ((new (...args: any[]) => string | number | import("../../cascader-panel/src/node").CascaderNodePathValue | (import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue)[]) | (() => import("../../cascader-panel/src/node").CascaderValue))[], unknown, unknown>>;
        readonly required: false;
        readonly validator: ((val: unknown) => boolean) | undefined;
        __epPropKey: true;
    };
    options: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderOption[]) | (() => import("element-plus/es/components/cascader-panel").CascaderOption[]))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderOption[], boolean>;
    props: import("element-plus/es/utils").EpPropFinalized<(new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps) | ((new (...args: any[]) => import("element-plus/es/components/cascader-panel").CascaderProps) | (() => import("element-plus/es/components/cascader-panel").CascaderProps))[], unknown, unknown, () => import("element-plus/es/components/cascader-panel").CascaderProps, boolean>;
}>> & {
    "onUpdate:modelValue"?: ((_: CascaderValue) => any) | undefined;
    onChange?: ((_: CascaderValue) => any) | undefined;
    onFocus?: ((evt: FocusEvent) => any) | undefined;
    onBlur?: ((evt: FocusEvent) => any) | undefined;
    onClear?: (() => any) | undefined;
    onVisibleChange?: ((val: boolean) => any) | undefined;
    onExpandChange?: ((val: CascaderValue) => any) | undefined;
    onRemoveTag?: ((val: import("../../cascader-panel/src/node").CascaderNodeValue | import("../../cascader-panel/src/node").CascaderNodePathValue) => any) | undefined;
}, {
    disabled: boolean;
    separator: string;
    props: import("element-plus/es/components/cascader-panel").CascaderProps;
    placement: import("element-plus/es/utils").EpPropMergeType<(new (...args: any[]) => "top" | "bottom" | "left" | "right" | "auto" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => import("element-plus/es/components/popper").Placement) | ((new (...args: any[]) => "top" | "bottom" | "left" | "right" | "auto" | "auto-start" | "auto-end" | "top-start" | "top-end" | "bottom-start" | "bottom-end" | "right-start" | "right-end" | "left-start" | "left-end") | (() => import("element-plus/es/components/popper").Placement))[], import("element-plus/es/components/popper").Placement, unknown>;
    options: import("element-plus/es/components/cascader-panel").CascaderOption[];
    valueOnClear: import("element-plus/es/utils").EpPropMergeType<readonly [StringConstructor, NumberConstructor, BooleanConstructor, FunctionConstructor], unknown, unknown>;
    validateEvent: import("element-plus/es/utils").EpPropMergeType<BooleanConstructor, unknown, unknown>;
    clearable: boolean;
    fallbackPlacements: import("element-plus/es/components/popper").Placement[];
    popperClass: string;
    teleported: import("element-plus/es/utils").EpPropMergeType<BooleanConstructor, unknown, unknown>;
    persistent: import("element-plus/es/utils").EpPropMergeType<BooleanConstructor, unknown, unknown>;
    debounce: number;
    filterable: boolean;
    filterMethod: (node: CascaderNode, keyword: string) => boolean;
    showAllLevels: import("element-plus/es/utils").EpPropMergeType<BooleanConstructor, unknown, unknown>;
    collapseTags: boolean;
    maxCollapseTags: number;
    collapseTagsTooltip: import("element-plus/es/utils").EpPropMergeType<BooleanConstructor, unknown, unknown>;
    beforeFilter: (value: string) => boolean | Promise<any>;
    tagType: import("element-plus/es/utils").EpPropMergeType<StringConstructor, "success" | "warning" | "info" | "primary" | "danger", unknown>;
    tagEffect: import("element-plus/es/utils").EpPropMergeType<StringConstructor, "dark" | "light" | "plain", unknown>;
}>;
declare const _default: __VLS_WithTemplateSlots<typeof __VLS_component, ReturnType<typeof __VLS_template>>;
export default _default;
type __VLS_WithTemplateSlots<T, S> = T & {
    new (): {
        $slots: S;
    };
};
