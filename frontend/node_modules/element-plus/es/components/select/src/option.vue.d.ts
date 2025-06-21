declare const _default: import("vue").DefineComponent<{
    /**
     * @description value of option
     */
    value: {
        required: true;
        type: (ObjectConstructor | NumberConstructor | StringConstructor | BooleanConstructor)[];
    };
    /**
     * @description label of option, same as `value` if omitted
     */
    label: (NumberConstructor | StringConstructor)[];
    created: BooleanConstructor;
    /**
     * @description whether option is disabled
     */
    disabled: BooleanConstructor;
}, {
    ns: {
        namespace: import("vue").ComputedRef<string>;
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
    id: import("vue").Ref<string>;
    containerKls: import("vue").ComputedRef<string[]>;
    currentLabel: import("vue").ComputedRef<any>;
    itemSelected: import("vue").ComputedRef<boolean>;
    isDisabled: import("vue").ComputedRef<any>;
    select: import("./token").SelectContext | undefined;
    hoverItem: () => void;
    updateOption: (query: string) => void;
    visible: import("vue").Ref<boolean>;
    hover: import("vue").Ref<boolean>;
    selectOptionClick: () => void;
    states: {
        index: number;
        groupDisabled: boolean;
        visible: boolean;
        hover: boolean;
    };
}, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, Record<string, any>, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    /**
     * @description value of option
     */
    value: {
        required: true;
        type: (ObjectConstructor | NumberConstructor | StringConstructor | BooleanConstructor)[];
    };
    /**
     * @description label of option, same as `value` if omitted
     */
    label: (NumberConstructor | StringConstructor)[];
    created: BooleanConstructor;
    /**
     * @description whether option is disabled
     */
    disabled: BooleanConstructor;
}>>, {
    disabled: boolean;
    created: boolean;
}>;
export default _default;
