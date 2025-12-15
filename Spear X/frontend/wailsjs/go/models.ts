export namespace main {
	
	export class Tool {
	    id: string;
	    name: string;
	    path: string;
	    fileName: string;
	    value: string;
	    command: string;
	    optional: string;
	    description: string;
	    tags: string[];
	    sourceUrl: string;
	    iconPath: string;
	    openCount: number;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    lastUsedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Tool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.path = source["path"];
	        this.fileName = source["fileName"];
	        this.value = source["value"];
	        this.command = source["command"];
	        this.optional = source["optional"];
	        this.description = source["description"];
	        this.tags = source["tags"];
	        this.sourceUrl = source["sourceUrl"];
	        this.iconPath = source["iconPath"];
	        this.openCount = source["openCount"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.lastUsedAt = this.convertValues(source["lastUsedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Category {
	    name: string;
	    icon: string;
	    tools: Tool[];
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.icon = source["icon"];
	        this.tools = this.convertValues(source["tools"], Tool);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Categories {
	    categories: Category[];
	
	    static createFrom(source: any = {}) {
	        return new Categories(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.categories = this.convertValues(source["categories"], Category);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class CleanupResult {
	    invalidToolsCount: number;
	    invalidCategoriesCount: number;
	    cleanedNotes: number;
	    migratedNotes: number;
	    invalidToolNames: string[];
	    migratedToolNames: string[];
	
	    static createFrom(source: any = {}) {
	        return new CleanupResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.invalidToolsCount = source["invalidToolsCount"];
	        this.invalidCategoriesCount = source["invalidCategoriesCount"];
	        this.cleanedNotes = source["cleanedNotes"];
	        this.migratedNotes = source["migratedNotes"];
	        this.invalidToolNames = source["invalidToolNames"];
	        this.migratedToolNames = source["migratedToolNames"];
	    }
	}
	export class FileInfo {
	    name: string;
	    isDir: boolean;
	    size: number;
	    modTime: string;
	    path: string;
	    extension: string;
	    isExecutable: boolean;
	
	    static createFrom(source: any = {}) {
	        return new FileInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.isDir = source["isDir"];
	        this.size = source["size"];
	        this.modTime = source["modTime"];
	        this.path = source["path"];
	        this.extension = source["extension"];
	        this.isExecutable = source["isExecutable"];
	    }
	}
	export class JavaConfig {
	    Java8: string;
	    Java11: string;
	    Java17: string;
	
	    static createFrom(source: any = {}) {
	        return new JavaConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Java8 = source["Java8"];
	        this.Java11 = source["Java11"];
	        this.Java17 = source["Java17"];
	    }
	}
	export class ScannedTool {
	    path: string;
	    category: string;
	    possibleFiles: string[];
	
	    static createFrom(source: any = {}) {
	        return new ScannedTool(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.category = source["category"];
	        this.possibleFiles = source["possibleFiles"];
	    }
	}

}

