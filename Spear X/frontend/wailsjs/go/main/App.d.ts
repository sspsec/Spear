// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {main} from '../models';

export function AddTool(arg1:main.Tool,arg2:string):Promise<void>;

export function DeleteCategory(arg1:string):Promise<void>;

export function DeleteTool(arg1:string,arg2:string):Promise<void>;

export function ExecuteCommand(arg1:string,arg2:string,arg3:string,arg4:string):Promise<void>;

export function GetCategories():Promise<main.Categories>;

export function GetFileInfo(arg1:string):Promise<{[key: string]: string}>;

export function GetFilePath(arg1:string):Promise<string>;

export function GetToolTypes():Promise<Array<string>>;

export function OpenDirectoryDialog():Promise<string>;

export function OpenFileDialog():Promise<{[key: string]: string}>;

export function OpenToolDirectory(arg1:string):Promise<void>;

export function SelectFile():Promise<string>;

export function UpdateCategoryTools(arg1:string,arg2:Array<main.Tool>):Promise<void>;

export function UpdateTool(arg1:string,arg2:string,arg3:main.Tool):Promise<void>;

export function UpdateToolDescription(arg1:string,arg2:string,arg3:string):Promise<void>;
