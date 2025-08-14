
<template>
  <div class="app-wrapper" :class="{ resizing: isResizing }">
    <!-- 左侧分类导航 -->
    <div class="sidebar" :class="{ resizing: isResizing }" :style="{ width: sidebarWidth + 'px' }">
      <div class="sidebar-header">
        <!-- SpearX 品牌标识 -->
        <div class="brand-logo" @click="openGitHub" title="访问 GitHub 项目页面">
          <!-- 简约呼吸光晕 -->
          <div class="elegant-glow"></div>
          <!-- SpearX 纯净文字 -->
          <div class="app-name">SpearX</div>
          <!-- 乔布斯式优雅签名 -->
          <div class="brand-signature">Created by Spe4r</div>
        </div>
      </div>
      
      <div class="category-nav">
        <div class="nav-section">
          <div class="nav-item all" @click="selectCategory('all')" :class="{active: selectedCategoryName === 'all'}">
            <span class="nav-icon">📦</span>
            <span class="nav-text">全部工具</span>
            <span class="nav-count">{{ getTotalToolCount() }}</span>
          </div>
        </div>

        <div class="nav-section">
          <draggable 
            v-model="sortableCategories" 
            @end="onCategorySortEnd"
            item-key="name"
            handle=".category-drag-handle"
            ghost-class="category-ghost"
            chosen-class="category-chosen"
            drag-class="category-drag"
          >
            <template #item="{ element: category }">
              <div 
                class="nav-item category-item" 
                @click="selectCategory(category.name || category.Name)"
                :class="{active: selectedCategoryName === (category.name || category.Name)}"
              >
                <span class="category-drag-handle" title="拖动排序">⋮⋮</span>
            <el-popover
              :visible="iconPopover.visible && iconPopover.categoryName === (category.name || category.Name)"
              placement="right"
              width="280px"
              trigger="manual"
              @hide="hideIconPopover"
              popper-class="dark-icon-popover"
            >
              <template #reference>
                <span 
                  class="nav-icon clickable-icon" 
                  @click.stop="showIconPopover(category.name || category.Name)"
                  :title="'点击更改图标'"
                >
                  {{ category.icon || getCategoryIcon(category.name || category.Name) }}
                </span>
              </template>
              <div class="icon-popover-content">
                <div class="icon-grid-compact">
                  <div 
                    v-for="icon in availableIcons" 
                    :key="icon"
                    class="icon-item-compact"
                    :class="{ active: iconPopover.selectedIcon === icon }"
                    @click="selectIconFromPopover(icon)"
                    :title="icon"
                  >
                    {{ icon }}
                  </div>
                </div>
                <div class="icon-popover-footer">
                  <div class="popover-buttons">
                    <el-button @click="hideIconPopover" size="small">取消</el-button>
                    <el-button type="primary" @click="updateCategoryIconFromPopover" size="small">确定</el-button>
                  </div>
                </div>
              </div>
            </el-popover>
            <span 
              v-if="editingCategory !== (category.name || category.Name)"
              class="nav-text editable-text" 
              @dblclick.stop="startInlineEditCategoryName(category.name || category.Name)"
              :title="'双击编辑分类名'"
            >
              {{ category.name || category.Name }}
            </span>
            <input 
              v-else
              v-model="editingCategoryName"
              class="nav-text-input"
              @blur="finishEditCategoryName"
              @keyup.enter="finishEditCategoryName"
              @keyup.esc="cancelEditCategoryName"
              ref="categoryNameInput"
            />
            <span class="nav-count">{{ (category.tools || category.Tool || []).length }}</span>
              <el-button 
              class="delete-category-btn"
                size="small" 
              type="danger"
              text
              @click.stop="deleteCategoryConfirm(category.name || category.Name)"
              title="删除分类"
            >
              <el-icon><Delete /></el-icon>
            </el-button>
              </div>
            </template>
          </draggable>
          

        </div>
      </div>
    </div>

    <!-- 可拖动分隔线 -->
    <div 
      class="sidebar-resizer" 
      :class="{ active: isResizing }"
      @mousedown="startResize"
      :style="{ left: sidebarWidth + 'px' }"
    ></div>

    <!-- 主内容区域 -->
    <div class="main-content" :class="{ resizing: isResizing }" :style="{ marginLeft: sidebarWidth + 'px' }">
      <div class="content-header">
        <div class="search-bar">
          <div class="search-wrapper">
            <el-input
              ref="searchInput"
              v-model="searchQuery"
              placeholder="搜索工具"
              class="ios-search"
              clearable
              @input="filterTools"
              @keyup.esc="clearSearch"
              @keyup.enter="executeFirstResult"
            >
              <template #prefix>
                <el-icon class="search-icon"><Search /></el-icon>
              </template>
            </el-input>
          </div>
          <div class="action-buttons">
            <el-button @click="showAddToolDialog" title="添加工具" class="icon-button">
              <el-icon><Plus /></el-icon>
            </el-button>
            <el-button @click="scanAndRefreshTools" title="扫描Resources目录" class="icon-button">
              <el-icon><Refresh /></el-icon>
            </el-button>
            <el-button @click="scanCustomDirectory" title="扫描自定义目录" class="icon-button">
              <el-icon><FolderOpened /></el-icon>
            </el-button>
            <el-button @click="showJavaConfigDialog" title="Java配置" class="icon-button">
              <el-icon><Setting /></el-icon>
            </el-button>
          </div>
            </div>
          </div>

      <!-- 工具网格 -->
      <div class="tools-container">
          <div class="current-category-title" v-if="selectedCategoryName !== 'all'">
            {{ selectedCategoryName }}
          </div>
          
          <div class="tools-grid">
            <draggable
            v-model="currentTools"
              :animation="150"
              ghost-class="ghost"
            @end="onDragEnd"
            item-key="name"
            class="tools-grid-inner"
          >
            <template #item="{ element: tool, index }">
              <div 
                  class="tool-card"
                  @click="executeTool(tool)"
              >
                <div class="tool-header">
                  <div class="tool-icon">
                    <span class="type-icon" v-html="getToolIcon(tool)"></span>
                  </div>
                  <div class="tool-title" :title="tool.name">{{ tool.name }}</div>
                </div>
                
                <div class="tool-body">
                  <div class="tool-tags">
                    <el-tag v-if="tool.value" size="small" :type="getTagType(tool.value)">{{ tool.value }}</el-tag>
                    <el-tag 
                      v-for="tag in (tool.tags || [])" 
                      :key="tag" 
                      size="small" 
                      type="primary"
                      @click="searchByTag(tag)"
                      class="clickable-tag"
                    >
                      {{ tag }}
                    </el-tag>
                  </div>
                </div>
                
                <div class="tool-footer">
                  <div class="action-group">
                    <div class="action-buttons">
                      <el-button 
                        size="small" 
                        @click.stop="openToolDirectoryByPath(tool.path)" 
                        text
                        title="打开目录"
                      >
                        <el-icon><Folder /></el-icon>
                      </el-button>
                      <el-button 
                        size="small" 
                        @click.stop="editTool(tool)" 
                        text
                        title="编辑工具"
                      >
                        <el-icon><Setting /></el-icon>
                      </el-button>
                      <el-button 
                        size="small" 
                        @click.stop="openToolNote(tool)" 
                        text
                        title="编辑笔记"
                      >
                        <el-icon><Document /></el-icon>
                      </el-button>
                      <el-button 
                        size="small" 
                        @click.stop="copyToolPath(tool)" 
                        text
                        title="复制路径"
                      >
                        <el-icon><CopyDocument /></el-icon>
                      </el-button>
                      <el-button 
                        size="small" 
                        type="danger" 
                        @click.stop="deleteToolConfirm(tool)" 
                        text
                        title="删除工具"
                      >
                        <el-icon><Delete /></el-icon>
                      </el-button>
                    </div>
                  </div>
                </div>
              </div>
              </template>
            </draggable>
          </div>
        </div>
      </div>
      








      <!-- 全局Java配置对话框 -->
      <el-dialog
        v-model="javaConfigDialog.visible"
        width="600px"
        :before-close="() => javaConfigDialog.visible = false"
        class="java-config-dialog mac-native-dialog"
      >
        <template #header>
          <div class="dialog-header-with-info">
            <span class="dialog-title">全局Java配置</span>
            <el-tooltip
              content="配置Java路径后，所有工具都将使用这些Java版本。可以手动输入完整路径，或点击输入框选择具体的Java可执行文件。如果不配置，将使用系统默认的java命令。"
              placement="bottom"
              :show-after="300"
              popper-class="java-config-tooltip"
            >
              <el-icon class="info-icon"><InfoFilled /></el-icon>
            </el-tooltip>
          </div>
        </template>
        <div class="java-config-content">
          
          <el-form label-width="80px" class="java-config-form">
            <el-form-item label="Java 8">
              <div class="java-path-input-group">
                <el-input 
                  v-model="javaConfigDialog.config.Java8" 
                  placeholder="留空使用系统默认，或点击输入框选择Java 8可执行文件路径"
                  @click="selectJavaPath('Java8')"
                  style="cursor: pointer;"
                  class="java-path-input"
                />
                <el-button 
                  @click.stop="clearJavaPath('Java8')" 
                  title="清除Java 8路径"
                  size="small"
                  type="danger"
                  class="java-clear-button"
                >
                  <el-icon><Delete /></el-icon>
                </el-button>
              </div>
            </el-form-item>
            
            <el-form-item label="Java 11">
              <div class="java-path-input-group">
                <el-input 
                  v-model="javaConfigDialog.config.Java11" 
                  placeholder="留空使用系统默认，或点击输入框选择Java 11可执行文件路径"
                  @click="selectJavaPath('Java11')"
                  style="cursor: pointer;"
                  class="java-path-input"
                />
                <el-button 
                  @click.stop="clearJavaPath('Java11')" 
                  title="清除Java 11路径"
                  size="small"
                  type="danger"
                  class="java-clear-button"
                >
                  <el-icon><Delete /></el-icon>
                </el-button>
              </div>
            </el-form-item>
            
            <el-form-item label="Java 17">
              <div class="java-path-input-group">
                <el-input 
                  v-model="javaConfigDialog.config.Java17" 
                  placeholder="留空使用系统默认，或点击输入框选择Java 17可执行文件路径"
                  @click="selectJavaPath('Java17')"
                  style="cursor: pointer;"
                  class="java-path-input"
                />
                <el-button 
                  @click.stop="clearJavaPath('Java17')" 
                  title="清除Java 17路径"
                  size="small"
                  type="danger"
                  class="java-clear-button"
                >
                  <el-icon><Delete /></el-icon>
                </el-button>
              </div>
            </el-form-item>
          </el-form>
        </div>
        
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="javaConfigDialog.visible = false">取消</el-button>
            <el-button type="primary" @click="saveJavaConfig">保存配置</el-button>
          </span>
        </template>
      </el-dialog>



      <!-- 命令输出区域 -->
      <el-card class="output-card" v-if="outputText">
        <template #header>
          <div class="output-header">
            <span>命令输出</span>
            <el-button :icon="Delete" @click="outputText = ''" title="清除输出" class="icon-button"></el-button>
          </div>
        </template>
        <pre class="output-content">{{ outputText }}</pre>
      </el-card>

      <!-- 右侧笔记编辑卡片 -->
      <div 
        v-if="noteDialog.visible" 
        class="note-editor-overlay"
        @click="closeNoteDialog"
      >
        <div 
          class="note-editor-card"
          @click.stop
        >
          <div class="note-card-header">
            <div class="note-card-title">
              <el-icon class="title-icon"><Document /></el-icon>
              <span>{{ (noteDialog.tool?.name || noteDialog.tool?.Name || '工具') + ' - 笔记' }}</span>
            </div>
            <div class="note-card-actions">
              <el-button 
                type="text" 
                @click="togglePreview"
                class="preview-toggle"
                size="small"
              >
                <el-icon><Edit v-if="noteDialog.isPreview" /><View v-else /></el-icon>
                {{ noteDialog.isPreview ? '编辑' : '预览' }}
              </el-button>
              <el-button 
                type="text" 
                @click="closeNoteDialog"
                :icon="Close"
                class="close-button"
                size="small"
              />
            </div>
          </div>
          
          <div class="note-card-content">
            <div v-if="!noteDialog.isPreview" class="note-editor-container">
              <el-input
                v-model="noteDialog.content"
                type="textarea"
                placeholder="在此输入Markdown格式的笔记..."
                :rows="25"
                resize="none"
                class="note-editor-textarea"
                @keydown.enter.stop
                @keyup.enter.stop
              />
            </div>
            
            <div v-else class="note-preview-container">
              <div class="markdown-preview" v-html="renderMarkdown(noteDialog.content || '')"></div>
            </div>
          </div>
          
          <div class="note-card-footer">
            <div class="note-info">
              <span class="note-tool-info">{{ noteDialog.tool?.categoryName || '未分类' }}</span>
            </div>
            <div class="note-actions">
              <el-button @click="closeNoteDialog" size="small">取消</el-button>
              <el-button type="primary" @click="saveNoteDialog" size="small">保存</el-button>
            </div>
          </div>
        </div>
      </div>

      <!-- 修改右键菜单实现 -->
      <div 
        v-if="contextMenu.visible" 
        class="context-menu"
        :style="{
          left: `${contextMenu.x}px`,
          top: `${contextMenu.y}px`
        }"
      >
        <div class="menu-item" @click="openToolDirectory">
          <el-icon><Folder /></el-icon>
          打开目录
        </div>
        <div class="menu-item" @click="showEditDialog">
          <el-icon><Edit /></el-icon>
          修改工具
        </div>
        <div class="menu-item delete" @click="deleteSelectedTool">
          <el-icon><Delete /></el-icon>
          删除工具
        </div>
      </div>

      <!-- 右侧编辑工具卡片 -->
      <div 
        v-if="editDialog.visible" 
        class="edit-tool-overlay"
        @click="closeEditDialog"
      >
        <div 
          class="edit-tool-card"
          @click.stop
        >
          <div class="edit-card-header">
            <div class="edit-card-title">
              <el-icon class="title-icon"><Tools /></el-icon>
              <span>{{ editDialog.isAddMode ? '添加工具' : '编辑工具' }}</span>
            </div>
            <el-button 
              type="text" 
              @click="closeEditDialog"
              :icon="Close"
              class="close-btn"
            />
          </div>
        <el-form :model="editDialog.tool" label-width="80px">
          <el-form-item label="工具名称">
            <el-input 
              v-model="editDialog.tool.name" 
              placeholder="请输入工具名称"
              @keydown.enter.stop
              @keyup.enter.stop
            />
          </el-form-item>
          
          <!-- 浏览器方式显示URL输入框 -->
          <el-form-item label="网页URL" v-if="editDialog.tool.value === 'Browser'">
            <el-input 
              v-model="editDialog.tool.url" 
              placeholder="请输入网页地址，如：https://www.example.com"
            >
              <template #prepend>
                <el-icon><Link /></el-icon>
              </template>
            </el-input>
          </el-form-item>
          
          <!-- 非浏览器方式显示工具路径 -->
          <el-form-item label="工具路径" v-if="editDialog.tool.value !== 'Browser'">
            <el-input 
              :value="getDisplayPath(editDialog.tool.path)" 
              placeholder="请点击右侧按钮选择目录，或使用下方文件浏览器"
              readonly
              @keydown.enter.stop
              @keyup.enter.stop
            >
              <template #prepend>
                <el-icon><Folder /></el-icon>
              </template>
              <template #append>
                <el-button 
                  @click.stop="selectToolDirectory" 
                  :icon="FolderOpened"
                  title="选择工具目录"
                >
                  打开
                </el-button>
              </template>
            </el-input>
          </el-form-item>
          
          <!-- 非浏览器方式显示文件名 -->
          <el-form-item label="文件名" v-if="editDialog.tool.value !== 'Browser'">
            <el-input 
              v-model="editDialog.tool.fileName" 
              placeholder="请点击右侧按钮选择文件，或点击下方文件浏览器中的文件"
              readonly
              @keydown.enter.stop
              @keyup.enter.stop
            >
              <template #prepend>
                <el-icon><Document /></el-icon>
              </template>
              <template #append>
                <el-button 
                  @click.stop="selectToolFile" 
                  :icon="FolderOpened"
                  title="选择文件"
                >
                  打开
                </el-button>
                <el-button 
                  @click.stop="clearFileSelection" 
                  :icon="Delete"
                  title="清除选择"
                  v-if="editDialog.tool.fileName"
                />
              </template>
            </el-input>
          </el-form-item>
          
          <el-form-item label="执行方式">
            <el-select 
              v-model="editDialog.tool.value" 
              placeholder="请选择执行方式"
              @change="onExecutionTypeChange"
            >
              <el-option
                v-for="javaType in ['Java8', 'Java11', 'Java17']"
                :key="javaType"
                :value="javaType"
              >
                <div class="execution-option">
                  <el-icon class="option-icon java-icon"><Coffee /></el-icon>
                  <div class="option-content">
                    <div class="option-title">{{ javaType }}</div>
                  </div>
                </div>
              </el-option>
              
              <el-option value="Open">
                <div class="execution-option">
                  <el-icon class="option-icon open-icon"><FolderOpened /></el-icon>
                  <div class="option-content">
                    <div class="option-title">系统打开</div>
                  </div>
                </div>
              </el-option>
              
              <el-option value="openterm">
                <div class="execution-option">
                  <el-icon class="option-icon terminal-icon"><Monitor /></el-icon>
                  <div class="option-content">
                    <div class="option-title">终端打开</div>
                  </div>
                </div>
              </el-option>
              
              <el-option value="Browser">
                <div class="execution-option">
                  <el-icon class="option-icon browser-icon"><Link /></el-icon>
                  <div class="option-content">
                    <div class="option-title">浏览器打开</div>
                  </div>
                </div>
              </el-option>
              

            </el-select>
          </el-form-item>
          
          <!-- 文件浏览器 -->

          
          <!-- 非浏览器、非Java方式显示命令输入框 -->
          <el-form-item label="命令" v-if="editDialog.tool.value !== 'Browser' && editDialog.tool.value !== 'Open' && !isJavaType(editDialog.tool.value)">
            <el-input 
              v-model="editDialog.tool.command" 
              :placeholder="editDialog.tool.value === 'openterm' ? '可选：自定义终端命令，留空则打开工具目录' : '可选：自定义命令'"
              :type="editDialog.tool.value === 'openterm' ? 'textarea' : 'text'"
              :rows="editDialog.tool.value === 'openterm' ? 2 : 1"
              @keydown.enter.stop
              @keyup.enter.stop
            />
            <div v-if="editDialog.tool.value === 'openterm'" class="form-tip">
              留空时默认打开工具目录；填写时可使用 {file} {filename} {path} 作为占位符
            </div>
          </el-form-item>
          
          <!-- 非浏览器、非终端方式显示可选参数 -->
          <el-form-item label="可选参数" v-if="editDialog.tool.value !== 'Browser' && editDialog.tool.value !== 'Open' && editDialog.tool.value !== 'openterm'">
            <el-input 
              v-model="editDialog.tool.optional" 
              placeholder="可选：命令行参数"
              @keydown.enter.stop
              @keyup.enter.stop
            />
          </el-form-item>
          
          <el-form-item label="工具描述">
            <el-input 
              v-model="editDialog.tool.description" 
              type="textarea" 
              :rows="3"
              placeholder="请输入工具描述（可选）"
              @keydown.enter.stop
              @keyup.enter.stop
            />
          </el-form-item>
          
          <el-form-item label="工具标签">
            <el-select
              v-model="editDialog.tool.tags"
              multiple
              filterable
              allow-create
              default-first-option
              placeholder="选择或创建标签"
              style="width: 100%"
            >
              <el-option
                v-for="tag in allTags"
                :key="tag"
                :label="tag"
                :value="tag"
              />
            </el-select>
          </el-form-item>
          
          <el-form-item label="所属分类">
            <el-select 
              v-model="editDialog.category"
              placeholder="请选择分类"
              allow-create
              filterable
              default-first-option
            >
              <el-option
                v-for="category in (categories.categories || categories.Category || [])"
                :key="category.name || category.Name"
                :label="category.name || category.Name"
                :value="category.name || category.Name"
              />
            </el-select>
          </el-form-item>
        </el-form>
        
        <div class="edit-card-footer">
          <el-button @click="closeEditDialog" size="large">取消</el-button>
          <el-button type="primary" @click="handleSubmit" size="large">
            {{ editDialog.isAddMode ? '添加工具' : '保存更改' }}
          </el-button>
    </div>
      </div>
    </div>

  </div>
</template>

<script>
import { ref, reactive, onMounted, onBeforeUnmount, nextTick, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { 
  Folder, 
  Edit, 
  Delete, 
  Plus,
  CopyDocument,
  Refresh,
  DocumentDelete,
  Select,
  Star,
  CircleCheckFilled,
  Close,
  View,
  Hide,
  Coffee,
  Monitor,
  Setting,
  Tools,
  Operation,
  Menu,
  MoreFilled,
  FolderOpened,
  Document,
  InfoFilled
} from '@element-plus/icons-vue'
import draggable from 'vuedraggable'

export default {
  name: 'App',
  components: {
    draggable,
    Folder,
    Edit,
    Delete,
    Plus,
    CopyDocument,
    Refresh,
    DocumentDelete,
    Select,
    Star,
    CircleCheckFilled,
    Close,
    View,
    Hide,
    Coffee,
    Monitor,
    Setting,
    Tools,
    Operation,
    Menu,
    MoreFilled,
    FolderOpened,
    Document,
    InfoFilled
  },
  setup() {
    const categories = ref({ Category: [] });
    const filteredCategories = ref([]);
    const sortableCategories = ref([]);
    const searchQuery = ref('');
    const showAddDialog = ref(false);
    const activeCategories = ref([]);
    const selectedCategoryName = ref('all');
    const currentTools = ref([]);
    const allTags = ref([]);
    
    // 分类编辑相关
    const editingCategory = ref(null);
    const editingCategoryName = ref('');
    
    // 图标气泡框相关
    const iconPopover = reactive({
      visible: false,
      categoryName: '',
      currentIcon: '',
      selectedIcon: '',

    });
    
    // 侧边栏拖动相关
    const sidebarWidth = ref(280);
    const isResizing = ref(false);
    

    const newTool = reactive({
      Name: '',
      Path: '',
      FileName: '',
      Value: '',
      Command: '',
      Optional: '',
      Description: '' // 添加描述字段
    });
    const selectedCategory = ref('');
    const outputText = ref('');
    const contextMenu = reactive({
      visible: false,
      x: 0,
      y: 0,
      selectedTool: null,
      selectedCategory: null
    });
    const editDialog = reactive({
      visible: false,
      isAddMode: false, // 添加模式标记
      tool: {
        name: '',
        path: '',
        fileName: '',
        value: '',
        command: '',
        optional: '',
        description: '',
        tags: [] // 添加标签字段
      },
      category: '',
      originalName: ''
    });





    // 可用图标列表
    const availableIcons = ref([
      // 文件夹和工具
      '📁', '📂', '📦', '🗂️', '📋', '📊', '📈', '📉', '📄', '📃', '📜', '📋', '🗃️', '🗄️',
      // 工具和设备
      '🔧', '🛠️', '⚙️', '🔩', '🔨', '🪛', '🔪', '🪓', '⛏️', '🛡️', '⚔️', '🗡️', '🏹', '🎯',
      // 计算机和技术
      '💻', '🖥️', '📱', '⌨️', '🖱️', '💾', '💿', '📀', '💽', '🖨️', '📺', '📟', '☎️', '📞',
      // 网络和全球
      '🌐', '🌍', '🌎', '🌏', '📡', '🛰️', '📶', '🔗', '🌐', '💫', '🌟', '⭐', '✨',
      // 科学和研究
      '🔍', '🔎', '🔬', '🧪', '⚗️', '🧬', '🔭', '🧮', '📐', '📏', '🧲', '⚖️', '🎓', '📚',
      // 航天和探索
      '🚀', '🛸', '🌌', '🪐', '🌠', '☄️', '🛰️', '🚁', '✈️', '🛩️', '🎈', '🪂',
      // 天气和自然
      '☀️', '🌙', '⭐', '🌟', '🔥', '⚡', '🌈', '☁️', '🌤️', '⛅', '🌦️', '🌧️', '⛈️', '🌩️',
      '❄️', '☃️', '⛄', '🌨️', '💧', '💦', '🌊', '🏔️', '🗻', '🌋', '🏕️', '🏞️',
      // 植物和生物
      '🌲', '🌳', '🌴', '🌱', '🌿', '🍀', '🌾', '🌵', '🌸', '🌺', '🌻', '🌷', '💐', '🌹',
      '🐚', '🦀', '🦞', '🐙', '🦑', '🐠', '🐟', '🐡', '🦈', '🐋', '🐬', '🦭', '🐧', '🦩',
      // 安全和加密
      '🔐', '🔒', '🔓', '🔑', '🗝️', '🛡️', '🔏', '🎫', '🏷️', '📍', '📌', '📎', '🔖', '🏴',
      // 娱乐和活动
      '🎮', '🕹️', '🎲', '♠️', '♥️', '♦️', '♣️', '🃏', '🎰', '🎯', '🎪', '🎨', '🖌️', '🖍️',
      // 交通工具
      '🚗', '🚕', '🚙', '🚌', '🚎', '🏎️', '🚓', '🚑', '🚒', '🚐', '🛻', '🚚', '🚛', '🚜',
      // 食物和饮品
      '🍎', '🍊', '🍋', '🍌', '🍉', '🍇', '🍓', '🫐', '🍈', '🍒', '🍑', '🥭', '🍍', '🥥',
      // 符号和形状
      '❤️', '🧡', '💛', '💚', '💙', '💜', '🖤', '🤍', '🤎', '💔', '❣️', '💕', '💖', '💗',
      '⭕', '❌', '⭐', '✅', '❎', '💯', '🔞', '📵', '🚯', '🚱', '🚳', '🚷', '🚸', '⚠️',
      // 箭头和方向
      '⬆️', '↗️', '➡️', '↘️', '⬇️', '↙️', '⬅️', '↖️', '↕️', '↔️', '↩️', '↪️', '⤴️', '⤵️',
      // 数字和标识
      '🔢', '1️⃣', '2️⃣', '3️⃣', '4️⃣', '5️⃣', '6️⃣', '7️⃣', '8️⃣', '9️⃣', '🔟', '💠', '🔷', '🔸'
    ]);
    
    // Java配置对话框
    const javaConfigDialog = reactive({
      visible: false,
      showDescription: false, // 控制说明文字的显示
      config: {
        Java8: '',
        Java11: '',
        Java17: ''
      }
    });

    // 笔记编辑对话框
    const noteDialog = reactive({
      visible: false,
      tool: null,
      content: '',
      isPreview: true, // 是否预览模式，默认为预览模式
      toolPath: '', // 工具路径
      toolName: '' // 工具名称
    });


    
    const toolTypes = ref([]);
    const searchInput = ref(null);

    const silentUpdate = ref(false);
    
    // 文件浏览器相关
    const fileBrowser = reactive({
      visible: false,
      currentPath: '',
      files: [],
      loading: false,
      breadcrumbs: []
    });
    
    // 提示相关变量






    // 加载分类和工具列表
    const loadCategories = async () => {
      try {
        const result = await window.go.main.App.GetCategories();
        categories.value = result;
        filteredCategories.value = result.categories || result.Category || [];
        sortableCategories.value = [...(result.categories || result.Category || [])];
        updateCurrentTools(); // 更新当前显示的工具
      } catch (err) {
        ElMessage.error(`加载工具列表失败: ${err}`);
      }
    };

    // 加载所有标签
    const loadAllTags = async () => {
      try {
        const tags = await window.go.main.App.GetAllTags();
        allTags.value = tags;
      } catch (err) {
        console.error('加载标签失败:', err);
      }
    };

    // 过滤工具
    const filterTools = () => {
      updateCurrentTools(); // 使用新的更新方法
    };

    // 执行工具
    const executeTool = async (tool) => {
      try {
        // 调试信息
        console.log('执行工具:', tool);
        console.log('工具类型:', tool.value || tool.Value);
        console.log('工具路径:', tool.path || tool.Path);
        console.log('工具完整信息:', JSON.stringify(tool, null, 2));
        
        const toolValue = tool.value || tool.Value;
        const toolPath = tool.path || tool.Path;
        const toolOptional = tool.optional || tool.Optional || '';
        const toolFileName = tool.fileName || tool.FileName || '';
        const toolCommand = tool.command || tool.Command || '';
        
        if (toolValue === 'custom' || (toolValue === 'openterm' && toolCommand)) {
          // 使用自定义命令（包括openterm有自定义命令的情况）
          await window.go.main.App.ExecuteCustomCommand(
            toolPath,
            toolOptional,
            toolValue,
            toolFileName,
            toolCommand
          );
        } else {
          // 使用预定义命令
        await window.go.main.App.ExecuteCommand(
            toolPath,
            toolOptional,
            toolValue,
            toolFileName
          );
        }
      } catch (err) {
        ElMessage.error(`执行失败: ${err}`);
      }
    };

    // 加载工具类型
    const loadToolTypes = async () => {
      try {
        toolTypes.value = await window.go.main.App.GetToolTypes();
      } catch (err) {
        ElMessage.error(`加载工具类型失败: ${err}`);
      }
    };

    // 重置新工具表单
    const resetNewToolForm = () => {
      Object.assign(newTool, {
        Name: '',
        Path: '',
        FileName: '',
        Value: '',
        Command: '',
        Optional: '',
        Description: '' // 重置描述字段
      });
      selectedCategory.value = '';
    };

    // 显示添加工具对话框（使用侧边卡片）
    const showAddToolDialog = () => {
      // 重置编辑对话框为添加模式
      Object.assign(editDialog, {
        visible: true,
        isAddMode: true, // 标记为添加模式
        tool: {
          name: '',
          path: '',
          fileName: '',
          value: '',
          command: '',
          optional: '',
          description: '',
          tags: []
        },
        category: '',
        originalName: ''
      });
    };

    // 扫描并刷新工具
    const scanAndRefreshTools = async () => {
      try {

        
        // 先清理无效路径并获取清理结果
        const cleanupResult = await window.go.main.App.CleanInvalidPaths();
        
        // 扫描所有工具
        const allScannedTools = await window.go.main.App.ScanResourcesForTools();
        
        // 过滤出真正的新工具
        const newTools = await window.go.main.App.GetNewToolsFromScanned(allScannedTools);
        
        // 重新加载工具列表以反映清理结果
        await loadCategories();
        await loadAllTags();
        
        // 清理结果记录到控制台，不显示用户提示
        if (cleanupResult.invalidToolsCount > 0) {
          const logMessage = `清理完成 - 无效工具: ${cleanupResult.invalidToolsCount} 个, 无效分类: ${cleanupResult.invalidCategoriesCount} 个, 清理笔记: ${cleanupResult.cleanedNotes} 个, 迁移笔记: ${cleanupResult.migratedNotes} 个`;
          console.log(logMessage);
          
          // 如果有笔记迁移，显示友好提示
          if (cleanupResult.migratedNotes > 0) {
            ElMessage({
              message: `已智能保护 ${cleanupResult.migratedNotes} 个工具笔记，避免因路径变更而丢失`,
              type: 'success',
              duration: 3000,
              showClose: true
            });
          }
        }
        
        if (newTools && newTools.length > 0) {
          // 显示扫描结果确认对话框
          const confirmResult = await ElMessageBox.confirm(
            `发现 ${newTools.length} 个新工具，是否自动添加到配置中？`,
            '扫描结果',
            {
              confirmButtonText: '自动添加',
              cancelButtonText: '取消',
              type: 'info',
              showClose: false,
              customClass: 'elegant-confirm-dialog'
            }
          );

          if (confirmResult === 'confirm') {
            // 自动添加新工具
            await window.go.main.App.AutoAddScannedTools(newTools);
            ElMessage.success(`成功添加 ${newTools.length} 个新工具`);
            
            // 重新加载工具列表
            await loadCategories();
            await loadAllTags();
          }
        } else {
          if (cleanupResult.invalidToolsCount === 0) {
            ElMessage.success('扫描完成，未发现新工具和无效路径');
          } else {
            ElMessage.success('扫描完成，未发现新工具');
          }
        }
      } catch (err) {
        ElMessage.error(`扫描工具失败: ${err}`);
      }
    };

    // 扫描自定义目录
    const scanCustomDirectory = async () => {
      try {
        // 选择目录
        const selectedPath = await window.go.main.App.SelectDirectory();
        if (!selectedPath) {
          return; // 用户取消选择
        }

        // 扫描自定义目录
        const allScannedTools = await window.go.main.App.ScanCustomDirectoryForTools(selectedPath);
        
        // 过滤出真正的新工具
        const newTools = await window.go.main.App.GetNewToolsFromScanned(allScannedTools);

        if (newTools && newTools.length > 0) {
          // 显示扫描结果确认对话框
          const confirmResult = await ElMessageBox.confirm(
            `在 "${selectedPath}" 中发现 ${newTools.length} 个新工具，是否自动添加到配置中？`,
            '自定义目录扫描结果',
            {
              confirmButtonText: '自动添加',
              cancelButtonText: '取消',
              type: 'info',
              showClose: false,
              customClass: 'elegant-confirm-dialog'
            }
          );

          if (confirmResult === 'confirm') {
            // 自动添加新工具
            await window.go.main.App.AutoAddScannedTools(newTools);
            ElMessage.success(`成功从自定义目录添加 ${newTools.length} 个新工具`);
            
            // 重新加载工具列表
            await loadCategories();
            await loadAllTags();
          }
        } else {
          ElMessage.info(`扫描完成，在 "${selectedPath}" 中未发现新工具`);
        }
      } catch (err) {
        ElMessage.error(`扫描自定义目录失败: ${err}`);
      }
    };

    // 打开GitHub页面 (macOS)
    const openGitHub = async () => {
      try {
        // 调用后端函数使用macOS默认浏览器打开GitHub链接
        await window.go.main.App.OpenGitHubPage();
        console.log('GitHub页面已打开');
      } catch (err) {
        console.error('打开GitHub页面失败:', err);
        ElMessage.error(`打开GitHub页面失败: ${err}`);
      }
    };

    // 选择工具路径
    const selectToolPath = async () => {
      try {
        const relativePath = await window.go.main.App.SelectFile();
        if (relativePath) {
          const pathParts = relativePath.split('/');
          const fileName = pathParts[pathParts.length - 1];
          const toolPath = pathParts.slice(0, -1).join('/');
          
          Object.assign(newTool, {
            Path: toolPath,
            FileName: fileName,
            Name: fileName.replace(/\.[^/.]+$/, "") // 如果工具名称为空，使用文件名（不带扩展名）
          });
        }
      } catch (err) {
        ElMessage.error(`${err}`);
      }
    };

    // 显示工具右键菜单
    const showToolMenu = (event, tool, categoryName) => {
      event.preventDefault();
      event.stopPropagation();
      
      // 获取视口尺寸
      const viewportHeight = window.innerHeight;
      const viewportWidth = window.innerWidth;
      
      // 获取菜单的预估尺寸
      const menuHeight = 120; // 菜单高度
      const menuWidth = 150;  // 菜单宽度
      
      // 获取鼠标点击的位置（相对于视口）
      let x = event.clientX;
      let y = event.clientY;
      
      // 计算菜单在各个方向上是否有足够空间
      const spaceBelow = viewportHeight - y;
      const spaceRight = viewportWidth - x;
      
      // 根据可用空间调整位置
      if (spaceBelow < menuHeight) {
        // 如果下方空间不足，向上显示
        y = y - menuHeight;
      }
      
      if (spaceRight < menuWidth) {
        // 如果右侧空间不足，向左显示
        x = x - menuWidth;
      }
      
      // 确保不会超出视口边界
      x = Math.max(0, Math.min(x, viewportWidth - menuWidth));
      y = Math.max(0, Math.min(y, viewportHeight - menuHeight));
      
      // 添加滚动偏移量以修正位置
      x += window.scrollX;
      y += window.scrollY;
      
      Object.assign(contextMenu, {
        visible: true,
        x,
        y,
        selectedTool: tool,
        selectedCategory: categoryName
      });
    };
    
    // 关闭右键菜单
    const closeContextMenu = (event) => {
      if (!event.target.closest('.context-menu') && contextMenu.visible) {
        contextMenu.visible = false;
      }
    };

    // 打开工具目录
    const openToolDirectory = async () => {
      if (!contextMenu.selectedTool) return;
      
      try {
        await window.go.main.App.OpenToolDirectory(contextMenu.selectedTool.Path);
        // 成功后再关闭菜单
        contextMenu.visible = false;
      } catch (err) {
        ElMessage.error(`打开目录失败: ${err.message || err}`);
      }
    };

    // 删除选中的工具
    const deleteSelectedTool = async () => {
      if (!contextMenu.selectedTool) return;

      // 先关闭右键菜单
      contextMenu.visible = false;

      try {
        await ElMessageBox.confirm(
          `确定要删除工具 "${contextMenu.selectedTool.Name}" 吗？这将同时删除对应的笔记。`,
          '删除确认',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
            customClass: 'elegant-confirm-dialog'
          }
        );

        // 删除工具对应的笔记
        const toolPath = contextMenu.selectedTool.Path || '';
        const toolName = contextMenu.selectedTool.Name || '';
        try {
          await window.go.main.App.DeleteToolNote(toolPath, toolName);
        } catch (noteErr) {
          console.log(`删除笔记失败（可能不存在）: ${noteErr}`);
        }

        // 删除工具
        await window.go.main.App.DeleteTool(
          contextMenu.selectedTool.Name,
          contextMenu.selectedCategory
        );
        
        ElMessage.success('工具及笔记删除成功');
        await loadCategories();
      } catch (err) {
        if (err !== 'cancel') {
          ElMessage.error(`删除失败: ${err}`);
        }
      }
    };

    // 显示编辑工具对话框
    const showEditDialog = () => {
      if (!contextMenu.selectedTool) return;

      // 先隐藏右键菜单
      contextMenu.visible = false;

      // 使用深拷贝避免数据污染
      Object.assign(editDialog, {
        visible: true,
        tool: JSON.parse(JSON.stringify(contextMenu.selectedTool)),
        category: contextMenu.selectedCategory,
        originalName: contextMenu.selectedTool.name
      });
      
      // 加载当前工具目录的文件
      if (contextMenu.selectedTool.path) {
        loadFileBrowser();
      }
    };

    // 关闭编辑工具对话框
    const closeEditDialog = () => {
      editDialog.visible = false;
      editDialog.isAddMode = false; // 重置添加模式标记
      editDialog.tool = {};
      editDialog.category = '';
      editDialog.originalName = '';
      fileBrowser.visible = false;
      fileBrowser.files = [];
    };

    // 提交编辑后的工具
    const submitToolEdit = async () => {
      try {
        await window.go.main.App.UpdateTool(
          editDialog.originalName,
          editDialog.category,
          editDialog.tool
        );
        closeEditDialog();
        await loadCategories(); // 重新加载工具列表
        await loadAllTags(); // 重新加载标签列表
      } catch (err) {
        ElMessage.error(`修改工具失败: ${err}`);
      }
    };

    // 处理提交按钮点击
    const handleSubmit = async () => {
      if (editDialog.isAddMode) {
        await submitToolAdd();
      } else {
        await submitToolEdit();
      }
    };

    // 提交添加工具（使用侧边卡片）
    const submitToolAdd = async () => {
      try {
        // 验证必要字段
        if (!editDialog.tool.name || !editDialog.tool.name.trim()) {
          ElMessage.error('请输入工具名称');
          return;
        }
        
        if (!editDialog.category || !editDialog.category.trim()) {
          ElMessage.error('请选择分类');
          return;
        }
        
        if (!editDialog.tool.value || !editDialog.tool.value.trim()) {
          ElMessage.error('请选择执行方式');
          return;
        }
        
        // 浏览器方式验证URL，其他方式验证路径
        if (editDialog.tool.value === 'Browser') {
          if (!editDialog.tool.url || !editDialog.tool.url.trim()) {
            ElMessage.error('请输入网页URL');
            return;
          }
          // 简单的URL格式验证
          const urlPattern = /^https?:\/\/.+/i;
          if (!urlPattern.test(editDialog.tool.url.trim())) {
            ElMessage.error('请输入有效的网页URL（以http://或https://开头）');
            return;
          }
        } else {
          if (!editDialog.tool.path || !editDialog.tool.path.trim()) {
            ElMessage.error('请选择工具路径');
            return;
          }
        }
        
        // 将工具对象转换为与后端兼容的格式
        const toolToAdd = {
          Name: editDialog.tool.name.trim(),
          Path: editDialog.tool.value === 'Browser' ? (editDialog.tool.url || '').trim() : (editDialog.tool.path || '').trim(),
          FileName: editDialog.tool.value === 'Browser' ? '' : (editDialog.tool.fileName || ''),
          Value: editDialog.tool.value,
          Command: editDialog.tool.command || '',
          Optional: editDialog.tool.optional || '',
          Description: editDialog.tool.description || '',
          Tags: editDialog.tool.tags || []
        };
        
        await window.go.main.App.AddTool(toolToAdd, editDialog.category);
        closeEditDialog();
        await loadCategories();
        await loadAllTags();
        // 成功消息由后端事件 'tool-added' 触发显示，避免重复
      } catch (err) {
        ElMessage.error(`添加工具失败: ${err}`);
      }
    };

    // 添加工具
    const addTool = async () => {
      try {
        // 验证必填字段
        if (!newTool.Name || !newTool.Name.trim()) {
          ElMessage.error('请输入工具名称');
          return;
        }
        
        if (!newTool.Value || !newTool.Value.trim()) {
          ElMessage.error('请选择执行类型');
          return;
        }
        
        // 浏览器方式验证URL，其他方式验证路径
        if (newTool.Value === 'Browser') {
          if (!newTool.URL || !newTool.URL.trim()) {
            ElMessage.error('请输入网页URL');
            return;
          }
          // 简单的URL格式验证
          const urlPattern = /^https?:\/\/.+/i;
          if (!urlPattern.test(newTool.URL.trim())) {
            ElMessage.error('请输入有效的网页URL（以http://或https://开头）');
            return;
          }
        } else {
          if (!newTool.Path || !newTool.Path.trim()) {
            ElMessage.error('请输入工具路径');
            return;
          }
        }
        
        await window.go.main.App.AddTool(newTool, selectedCategory.value);

        await loadCategories();
      } catch (err) {
        ElMessage.error(`添加工具失败: ${err}`);
      }
    };

    // 选择编辑工具路径
    const selectEditToolPath = async () => {
      try {
        const relativePath = await window.go.main.App.SelectFile();
        if (relativePath) {
          const pathParts = relativePath.split('/');
          const fileName = pathParts[pathParts.length - 1];
          const toolPath = pathParts.slice(0, -1).join('/');
          
          Object.assign(editDialog.tool, {
            Path: toolPath,
            FileName: fileName,
          });
        }
      } catch (err) {
        ElMessage.error(`${err}`);
      }
    };

    // 确认删除分类
    const confirmDeleteCategory = async (categoryName) => {
      try {
        await ElMessageBox.confirm(
          `确定要删除分类 "${categoryName}" 吗？\n该分类下的所有工具也会被删除。`,
          '删除确认',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
            customClass: 'elegant-confirm-dialog'
          }
        );
        
        await deleteCategory(categoryName);
      } catch (err) {
        if (err !== 'cancel') {
          ElMessage.error(`删除失败: ${err}`);
        }
      }
    };

    // 显示Java配置对话框
    const showJavaConfigDialog = async () => {
      try {
        // 加载当前的Java配置
        const config = await window.go.main.App.GetJavaConfig();
        if (config) {
          Object.assign(javaConfigDialog.config, config);
        }
        javaConfigDialog.visible = true;
      } catch (err) {
        ElMessage.error(`加载Java配置失败: ${err}`);
        javaConfigDialog.visible = true; // 即使加载失败也允许用户配置
      }
    };

    // 选择Java路径
    const selectJavaPath = async (javaVersion) => {
      try {
        const selectedPath = await window.go.main.App.SelectJavaPath();
        if (selectedPath) {
          javaConfigDialog.config[javaVersion] = selectedPath;
        }
      } catch (err) {
        ElMessage.error(`选择Java路径失败: ${err}`);
      }
    };

    // 清除Java路径
    const clearJavaPath = (javaVersion) => {
      javaConfigDialog.config[javaVersion] = '';
    };

    // 保存Java配置
    const saveJavaConfig = async () => {
      try {
        await window.go.main.App.SaveJavaConfig(javaConfigDialog.config);
        ElMessage.success('Java配置保存成功');
        javaConfigDialog.visible = false;
      } catch (err) {
        ElMessage.error(`保存Java配置失败: ${err}`);
      }
    };



    // 调试所有工具路径
    const debugAllPaths = async () => {
      try {
        await window.go.main.App.DebugAllToolPaths();
        ElMessage.success('路径调试信息已输出到控制台');
      } catch (err) {
        ElMessage.error(`调试失败: ${err}`);
      }
    };

    // 清理和修复工具路径
    const cleanupPaths = async () => {
      try {
        const result = await ElMessageBox.confirm(
          '这将自动修复所有工具的路径配置，确定要继续吗？',
          '修复路径',
          {
            confirmButtonText: '确定修复',
            cancelButtonText: '取消',
            type: 'warning'
          }
        );

        if (result === 'confirm') {
          await window.go.main.App.CleanupToolPaths();
          ElMessage.success('路径修复完成，请查看控制台输出');
          await loadCategories(); // 重新加载配置
        }
      } catch (err) {
        if (err !== 'cancel') {
          ElMessage.error(`路径修复失败: ${err}`);
        }
      }
    };

    // 删除分类
    const deleteCategory = async (categoryName) => {
      try {
        await window.go.main.App.DeleteCategory(categoryName);
        ElMessage.success('删除成功');
        await loadCategories();
      } catch (err) {
        ElMessage.error(`删除失败: ${err}`);
      }
    };

    // 处理全局键盘事件
    const handleGlobalKeydown = (event) => {
      if (event.target.tagName === 'INPUT' || event.target.tagName === 'TEXTAREA') {
        // 处理输入框中的特殊键
        if (event.code === 'Escape') {
          clearSearch();
        } else if (event.code === 'Enter') {
          executeFirstResult();
        }
        return;
      }
      
      // 处理全局键盘事件
      if (event.code === 'Space') {
        event.preventDefault();
        handleSpaceSearch();
      } else if (event.code === 'Escape') {
        clearSearch();
      }
      // 移除重复的Enter处理，只在输入框中处理Enter
    };

    // 执行第一个搜索结果
    const executeFirstResult = () => {
      // 如果有搜索结果，从当前显示的工具列表中获取第一个
      if (currentTools.value.length > 0) {
        const firstTool = currentTools.value[0];
        executeTool(firstTool);
        // 可选：清除搜索
        clearSearch();
        return;
      }
      // 如果没有搜索结果，显示提示
      if (searchQuery.value) {
        ElMessage.warning('没有找到匹配的工具');
      }
    };

    // 清除搜索
    const clearSearch = () => {
      searchQuery.value = '';
      filterTools();
      // 如果搜索框有焦点，则移除焦点
      if (searchInput.value && document.activeElement === searchInput.value.$el.querySelector('input')) {
        searchInput.value.blur();
      }
    };

    // 处理空格搜索
    const handleSpaceSearch = () => {
      if (searchInput.value) {
        searchInput.value.focus();
      }
    };



    // 选择分类
    const selectCategory = (categoryName) => {
      selectedCategoryName.value = categoryName;
      updateCurrentTools();
    };

    // 更新当前显示的工具
    const updateCurrentTools = () => {
      const categoryList = categories.value.categories || categories.value.Category || [];
      
      if (selectedCategoryName.value === 'all') {
        // 显示所有工具
        let allTools = [];
        categoryList.forEach(category => {
          const tools = category.tools || category.Tool || [];
          allTools = allTools.concat(tools.map(tool => ({
            ...tool,
            categoryName: category.name || category.Name
          })));
        });
        currentTools.value = allTools;
      } else {
        // 显示特定分类的工具
        const selectedCategory = categoryList.find(cat => (cat.name || cat.Name) === selectedCategoryName.value);
        const tools = selectedCategory ? (selectedCategory.tools || selectedCategory.Tool || []) : [];
        currentTools.value = tools.map(tool => ({
          ...tool,
          categoryName: selectedCategory ? (selectedCategory.name || selectedCategory.Name) : ''
        }));
      }
      
      // 应用搜索过滤
      if (searchQuery.value) {
        applySearchFilter();
      }
    };

    // 应用搜索过滤
    const applySearchFilter = async () => {
      const query = searchQuery.value.toLowerCase();
      
      // 检查是否是标签搜索
      if (query.startsWith('标签:')) {
        const tagQuery = query.replace('标签:', '').trim();
        currentTools.value = currentTools.value.filter(tool => {
          if (!tool.tags) return false;
          return tool.tags.some(tag => tag.toLowerCase().includes(tagQuery));
        });
      } else {
        // 扩展搜索：工具名称、描述、标签、笔记内容
        const searchPromises = currentTools.value.map(async (tool) => {
          const nameMatch = tool.name.toLowerCase().includes(query);
          const pathMatch = tool.path ? tool.path.toLowerCase().includes(query) : false;
          const descMatch = tool.description ? tool.description.toLowerCase().includes(query) : false;
          const urlMatch = tool.sourceUrl ? tool.sourceUrl.toLowerCase().includes(query) : false;
          const tagMatch = tool.tags ? tool.tags.some(tag => tag.toLowerCase().includes(query)) : false;
          
          // 搜索笔记内容
          let noteMatch = false;
          try {
            const note = await window.go.main.App.GetToolNote(tool.name);
            if (note && note.toLowerCase().includes(query)) {
              noteMatch = true;
            }
          } catch (err) {
            // 忽略笔记获取错误
          }
          
          return {
            tool,
            matches: nameMatch || pathMatch || descMatch || urlMatch || tagMatch || noteMatch
          };
        });
        
        const results = await Promise.all(searchPromises);
        currentTools.value = results.filter(result => result.matches).map(result => result.tool);
      }
    };

    // 获取总工具数量
    const getTotalToolCount = () => {
      const categoryList = categories.value.categories || categories.value.Category || [];
      return categoryList.reduce((total, category) => {
        const tools = category.tools || category.Tool || [];
        return total + tools.length;
      }, 0);
    };



    // 根据分类名称获取图标
    const getCategoryIcon = (categoryName) => {
      const iconMap = {
        'java8': '☕',
        'java11': '☕', 
        'java17': '☕',
        'webshell': '🐚',
        'pentest': '🔍',
        'comprehensive': '🛠️',
        'blue-team': '🛡️',
        'databases': '🗄️',
        'framework': '🏗️',
        'info': 'ℹ️',
        'Intranet': '🌐',
        'other': '📦',
        'proxy': '🔀'
      };
      return iconMap[categoryName] || '📁';
    };

    // 根据工具类型获取图标
    const getToolIcon = (tool) => {
      // 如果是浏览器类型，尝试获取网站图标
      if (tool.value === 'Browser' || tool.Value === 'Browser') {
        return getWebsiteIcon(tool.path || tool.Path);
      }
      
      const toolType = tool.value || tool.Value;
      const iconMap = {
        'Java8': '☕',      // Java咖啡图标
        'Java11': '☕',     // Java咖啡图标
        'Java17': '☕',     // Java咖啡图标
        'Open': '📱',       // 打开文件图标
        'openterm': '🖥️',   // 终端图标
        'Browser': '🌐',    // 浏览器图标
        'custom': '⚙️'      // 自定义命令图标
      };
      return iconMap[toolType] || '🔧';
    };

    // 获取网站图标
    const getWebsiteIcon = (url) => {
      if (!url) return '🌐';
      
      try {
        const urlObj = new URL(url);
        const domain = urlObj.hostname;
        
        // 构建 favicon URL - 使用 Google 的 favicon 服务
        const faviconUrl = `https://www.google.com/s2/favicons?domain=${domain}&sz=32`;
        
        // 返回 HTML img 标签而不是 emoji
        return `<img src="${faviconUrl}" alt="favicon" class="website-favicon" onerror="this.style.display='none';this.nextSibling.style.display='inline';" /><span style="display:none;">🌐</span>`;
      } catch (error) {
        return '🌐';
      }
    };

    // 获取标签类型
    const getTagType = (value) => {
      const typeMap = {
        'Java8': 'warning',
        'Java11': 'warning', 
        'Java17': 'warning',
        'Open': 'success',
        'openterm': 'info'
      };
      return typeMap[value] || 'primary';
    };

    // 格式化日期
    const formatDate = (date) => {
      return date.toLocaleDateString('zh-CN');
    };

    // 编辑工具
    const editTool = (tool) => {
      // 设置右键菜单状态以便重用现有方法
      contextMenu.selectedTool = tool;
      contextMenu.selectedCategory = tool.categoryName;
      showEditDialog();  // 复用现有的编辑对话框方法
    };

    // 复制工具路径
    const copyToolPath = async (tool) => {
      try {
        // 调用后端获取绝对路径
        const toolPath = tool.path || tool.Path || '';
        const fileName = tool.fileName || tool.FileName || '';
        const absolutePath = await window.go.main.App.GetToolAbsolutePath(toolPath, fileName);
        
        await navigator.clipboard.writeText(absolutePath);
        ElMessage.success('绝对路径已复制到剪贴板');
      } catch (err) {
        // 如果现代API失败，尝试备用方案
        try {
          // 如果后端调用失败，使用相对路径作为备用
          const fallbackPath = (tool.path || tool.Path || '') + (tool.fileName || tool.FileName ? `/${tool.fileName || tool.FileName}` : '');
          
          const textArea = document.createElement('textarea');
          textArea.value = fallbackPath;
          document.body.appendChild(textArea);
          textArea.select();
          document.execCommand('copy');
          document.body.removeChild(textArea);
          ElMessage.success('路径已复制到剪贴板');
        } catch (fallbackErr) {
          ElMessage.error('复制失败，请手动复制路径');
          console.error('复制失败:', err, fallbackErr);
        }
      }
    };

    // 确认删除工具
    const deleteToolConfirm = async (tool) => {
      try {
        await ElMessageBox.confirm(
          `确定要删除工具 "${tool.name}" 吗？这将同时删除对应的笔记。`,
          '删除确认',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
            customClass: 'elegant-confirm-dialog'
          }
        );

        // 找到工具所属的分类名称
        let categoryName = '';
        for (const category of filteredCategories.value) {
          const foundTool = category.tools?.find(t => t.name === tool.name) || 
                           category.Tool?.find(t => t.Name === tool.name);
          if (foundTool) {
            categoryName = category.name || category.Name;
            break;
          }
        }

        if (!categoryName) {
          throw new Error('无法找到工具所属分类');
        }

        // 删除工具对应的笔记
        const toolPath = tool.path || tool.Path || '';
        const toolName = tool.name || tool.Name || '';
        try {
          await window.go.main.App.DeleteToolNote(toolPath, toolName);
        } catch (noteErr) {
          console.log(`删除笔记失败（可能不存在）: ${noteErr}`);
        }

        // 删除工具
        await window.go.main.App.DeleteTool(tool.name, categoryName);
        ElMessage.success('工具及笔记删除成功');
        await loadCategories();
        updateCurrentTools();
      } catch (err) {
        if (err !== 'cancel') {
          ElMessage.error(`删除失败: ${err}`);
        }
      }
    };

    // 显示工具笔记
    const showToolNotes = (tool) => {
      ElMessage.info(`暂未实现笔记功能 - ${tool.Name}`);
    };

    // 通过路径打开工具目录
    const openToolDirectoryByPath = async (path) => {
      try {
        await window.go.main.App.OpenToolDirectory(path);
      } catch (err) {
        ElMessage.error(`打开目录失败: ${err.message || err}`);
      }
    };

    // 通过标签搜索
    const searchByTag = (tag) => {
      searchQuery.value = `标签:${tag}`;
      filterTools();
    };

    // 直接添加未命名分类
    const showAddCategoryDialog = async () => {
      try {
        // 生成唯一的分类名称
        let categoryName = '未命名';
        let counter = 0;
        
        // 获取当前所有分类名称
        const existingNames = new Set();
        const categoryList = categories.value.categories || categories.value.Category || [];
        categoryList.forEach(category => {
          existingNames.add(category.name || category.Name);
        });
        
        // 如果"未命名"已存在，则尝试"未命名1"、"未命名2"等
        while (existingNames.has(categoryName)) {
          counter++;
          categoryName = `未命名${counter}`;
        }
        
        await window.go.main.App.AddCategory(categoryName);
        ElMessage.success('分类添加成功，双击可重命名');
        await loadCategories();
      } catch (err) {
        ElMessage.error(`添加分类失败: ${err}`);
      }
    };

    // 添加新分类（保留用于兼容性）
    const addNewCategory = async () => {
      // 这个函数现在不再使用，但保留以防其他地方有引用
      await showAddCategoryDialog();
    };

    // 开始内联编辑分类名称
    const startInlineEditCategoryName = (categoryName) => {
      editingCategory.value = categoryName;
      editingCategoryName.value = categoryName;
      nextTick(() => {
        const input = document.querySelector('.nav-text-input');
        if (input) {
          input.focus();
          input.select();
        }
      });
    };

    // 完成编辑分类名称
    const finishEditCategoryName = async () => {
      if (!editingCategoryName.value.trim()) {
        ElMessage.error('分类名称不能为空');
        return;
      }
      
      const oldName = editingCategory.value;
      const newName = editingCategoryName.value.trim();
      
      if (oldName !== newName) {
        try {
          await updateCategoryName(oldName, newName);
        } catch (err) {
          ElMessage.error(`修改分类名称失败: ${err}`);
        }
      }
      
      cancelEditCategoryName();
    };

    // 取消编辑分类名称
    const cancelEditCategoryName = () => {
      editingCategory.value = null;
      editingCategoryName.value = '';
    };

    // 更新分类名称
    const updateCategoryName = async (oldName, newName) => {
      if (oldName === newName) {
        return; // 名称未改变
      }

      try {
        await window.go.main.App.UpdateCategoryName(oldName, newName);
        ElMessage.success('分类名称更新成功');
        
        // 如果当前选中的是被更新的分类，更新选中状态
        if (selectedCategoryName.value === oldName) {
          selectedCategoryName.value = newName;
        }
        
        await loadCategories();
        updateCurrentTools();
      } catch (err) {
        ElMessage.error(`更新分类名称失败: ${err}`);
      }
    };

    // 显示图标气泡框
    const showIconPopover = (categoryName) => {
      const category = (categories.value.categories || categories.value.Category || [])
        .find(cat => (cat.name || cat.Name) === categoryName);
      
      iconPopover.categoryName = categoryName;
      iconPopover.currentIcon = category?.icon || getCategoryIcon(categoryName);
      iconPopover.selectedIcon = category?.icon || getCategoryIcon(categoryName);

      iconPopover.visible = true;
    };

    // 隐藏图标气泡框
    const hideIconPopover = () => {
      iconPopover.visible = false;
      iconPopover.categoryName = '';
    };

    // 从气泡框选择图标
    const selectIconFromPopover = (icon) => {
      iconPopover.selectedIcon = icon;

    };

    // 从气泡框更新分类图标
    const updateCategoryIconFromPopover = async () => {
      if (!iconPopover.selectedIcon) {
        ElMessage.warning('请选择一个图标');
        return;
      }

      try {
        await window.go.main.App.UpdateCategoryIcon(
          iconPopover.categoryName, 
          iconPopover.selectedIcon
        );
        ElMessage.success('分类图标更新成功');
        hideIconPopover();
        await loadCategories();
      } catch (err) {
        ElMessage.error(`更新分类图标失败: ${err}`);
      }
    };

    // 分类拖动排序结束处理
    const onCategorySortEnd = async () => {
      try {
        // 将sortableCategories的顺序同步到categories
        categories.value.Category = [...sortableCategories.value];
        categories.value.categories = [...sortableCategories.value];
        
        // 保存新的分类顺序
        await window.go.main.App.UpdateCategoriesOrder(sortableCategories.value);
        ElMessage.success('分类顺序已更新');
      } catch (err) {
        ElMessage.error(`更新分类顺序失败: ${err}`);
        // 如果保存失败，恢复原来的顺序
        await loadCategories();
      }
    };

    // 确认删除分类
    const deleteCategoryConfirm = async (categoryName) => {
      try {
        const category = (categories.value.categories || categories.value.Category || [])
          .find(cat => (cat.name || cat.Name) === categoryName);
        
        const toolCount = (category?.tools || category?.Tool || []).length;
        const message = toolCount > 0 
          ? `确定要删除分类 "${categoryName}" 吗？这将删除该分类下的 ${toolCount} 个工具及其配置。`
          : `确定要删除分类 "${categoryName}" 吗？`;

        await ElMessageBox.confirm(message, '删除确认', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning',
          dangerouslyUseHTMLString: false,
          customClass: 'elegant-confirm-dialog'
        });

        await window.go.main.App.DeleteCategory(categoryName);
        ElMessage.success('分类删除成功');
        
        // 如果当前选中的是被删除的分类，切换到全部工具
        if (selectedCategoryName.value === categoryName) {
          selectedCategoryName.value = 'all';
        }
        
        await loadCategories();
        updateCurrentTools();
      } catch (err) {
        if (err !== 'cancel') {
          ElMessage.error(`删除分类失败: ${err}`);
        }
      }
    };

    // 打开工具笔记
    const openToolNote = async (tool) => {
      try {
        const toolPath = tool.path || tool.Path || '';
        const toolName = tool.name || tool.Name || '';
        
        if (!toolPath || !toolName) {
          ElMessage.error('工具信息不完整，无法打开笔记');
          return;
        }
        
        // 使用新的API读取笔记（从工具文件夹中）
        let noteContent = await window.go.main.App.GetToolNote(toolPath, toolName);
        
        // 设置笔记对话框数据
        noteDialog.tool = tool;
        noteDialog.content = noteContent || '';
        noteDialog.toolPath = toolPath;
        noteDialog.toolName = toolName;
        noteDialog.isPreview = true;
        noteDialog.visible = true;
      } catch (err) {
        ElMessage.error(`加载笔记失败: ${err}`);
      }
    };

    // 注意：旧的笔记迁移逻辑已移至后端自动处理

    // 关闭笔记编辑器
    const closeNoteDialog = () => {
      noteDialog.visible = false;
      noteDialog.tool = null;
      noteDialog.content = '';
      noteDialog.toolPath = '';
      noteDialog.toolName = '';
      noteDialog.isPreview = false;
    };

    // 切换预览模式
    const togglePreview = () => {
      noteDialog.isPreview = !noteDialog.isPreview;
    };

    // 注意：新版本直接使用工具路径和名称，无需生成ID

    // 保存工具笔记
    const saveNoteDialog = async () => {
      try {
        if (!noteDialog.toolPath || !noteDialog.toolName) {
          ElMessage.error('工具信息不完整，无法保存笔记');
          return;
        }
        
        await window.go.main.App.SaveToolNote(noteDialog.toolPath, noteDialog.toolName, noteDialog.content);
        ElMessage.success('笔记保存成功');
        closeNoteDialog();
      } catch (err) {
        ElMessage.error(`保存笔记失败: ${err}`);
      }
    };

    // 获取工具的笔记预览（用于显示是否有笔记）
    const getToolNotePreview = async (tool) => {
      try {
        const toolPath = tool.path || tool.Path || '';
        const toolName = tool.name || tool.Name || '';
        
        if (!toolPath || !toolName) return '';
        
        const content = await window.go.main.App.GetToolNote(toolPath, toolName);
        return content ? content.substring(0, 100) + '...' : '';
      } catch (err) {
        return '';
      }
    };

    // 简单的Markdown渲染器
    const renderMarkdown = (content) => {
      if (!content) return '';
      
      let html = content
        // 标题
        .replace(/^### (.*$)/gm, '<h3>$1</h3>')
        .replace(/^## (.*$)/gm, '<h2>$1</h2>')
        .replace(/^# (.*$)/gm, '<h1>$1</h1>')
        // 粗体和斜体
        .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
        .replace(/\*(.*?)\*/g, '<em>$1</em>')
        // 代码块 - 增强版本，带复制按钮
        .replace(/```([\s\S]*?)```/g, (match, code) => {
          const codeId = 'code-' + Math.random().toString(36).substr(2, 9);
          const trimmedCode = code.trim();
          
          // 解析语言类型和代码内容
          const lines = trimmedCode.split('\n');
          let language = '';
          let codeContent = trimmedCode;
          
          // 检查第一行是否是语言标识
          if (lines.length > 1 && lines[0].match(/^[a-zA-Z0-9_-]+$/)) {
            language = lines[0];
            codeContent = lines.slice(1).join('\n');
          }
          
          // 保存纯代码内容用于复制
          const encodedCode = btoa(encodeURIComponent(codeContent));
          
          return `<div class="code-block-container">
            <button class="code-copy-btn" onclick="copyCodeBlock('${codeId}')" title="复制代码" data-code="${encodedCode}">复制</button>
            <pre><code id="${codeId}" class="language-${language}">${trimmedCode}</code></pre>
          </div>`;
        })
        .replace(/`(.*?)`/g, '<code class="inline-code">$1</code>')
        // 链接
        .replace(/\[([^\]]+)\]\(([^)]+)\)/g, '<a href="$2" target="_blank">$1</a>')
        // 列表
        .replace(/^\* (.*$)/gm, '<li>$1</li>')
        .replace(/^- (.*$)/gm, '<li>$1</li>')
        // 换行
        .replace(/\n/g, '<br/>');
      
      // 处理列表包装
      html = html.replace(/(<li>.*<\/li>)/g, '<ul>$1</ul>');
      
      return html;
    };

    // 全局代码块复制函数
    const copyCodeBlock = async (codeId) => {
      try {
        // 优先从按钮的data-code属性获取纯代码内容
        const copyBtn = document.querySelector(`[onclick="copyCodeBlock('${codeId}')"]`);
        let codeText = '';
        
        if (copyBtn && copyBtn.dataset.code) {
          // 从data属性解码纯代码内容（不包含语言标识）
          codeText = decodeURIComponent(atob(copyBtn.dataset.code));
        } else {
          // 备用方案：从DOM元素获取
          const codeElement = document.getElementById(codeId);
          if (!codeElement) {
            console.error('Code element not found:', codeId);
            return;
          }
          codeText = codeElement.textContent || codeElement.innerText;
        }
        
        await navigator.clipboard.writeText(codeText);
        ElMessage.success('代码已复制到剪贴板');
      } catch (err) {
        // 备用方案
        try {
          const copyBtn = document.querySelector(`[onclick="copyCodeBlock('${codeId}')"]`);
          let codeText = '';
          
          if (copyBtn && copyBtn.dataset.code) {
            codeText = decodeURIComponent(atob(copyBtn.dataset.code));
          } else {
            const codeElement = document.getElementById(codeId);
            codeText = codeElement.textContent || codeElement.innerText;
          }
          
          const textArea = document.createElement('textarea');
          textArea.value = codeText;
          document.body.appendChild(textArea);
          textArea.select();
          document.execCommand('copy');
          document.body.removeChild(textArea);
          ElMessage.success('代码已复制到剪贴板');
        } catch (fallbackErr) {
          ElMessage.error('复制失败，请手动复制');
          console.error('Copy failed:', err, fallbackErr);
        }
      }
    };

    // 将复制函数绑定到全局window对象，以便HTML中的onclick可以调用
    window.copyCodeBlock = copyCodeBlock;

    // 拖拽结束
    const onDragEnd = async () => {
      try {
        silentUpdate.value = true;
        // 这里需要根据当前选择的分类来更新
        if (selectedCategoryName.value !== 'all') {
          const categoryList = categories.value.categories || categories.value.Category || [];
          const category = categoryList.find(cat => (cat.name || cat.Name) === selectedCategoryName.value);
          if (category) {
            const categoryName = category.name || category.Name;
            await window.go.main.App.UpdateCategoryTools(categoryName, currentTools.value);
        ElMessage.success('工具顺序已更新');
          }
        }
        await loadCategories();
        updateCurrentTools();
      } catch (err) {
        ElMessage.error(`更新工具顺序失败: ${err}`);
      } finally {
        silentUpdate.value = false;
      }
    };

    // 文件浏览器相关方法
    const getExecutionTypeLabel = (type) => {
      const labels = {
        'Java8': 'Java 8',
        'Java11': 'Java 11', 
        'Java17': 'Java 17',
        'Open': '系统打开',
        'openterm': '终端打开'
      };
      return labels[type] || type;
    };

    // 判断是否为Java执行类型
    const isJavaType = (type) => {
      return ['Java8', 'Java11', 'Java17'].includes(type);
    };

    const onExecutionTypeChange = () => {
      if (editDialog.tool.value && editDialog.tool.value !== 'custom' && editDialog.tool.path) {
        loadFileBrowser();
      }
    };

    const toggleFileBrowser = () => {
      fileBrowser.visible = !fileBrowser.visible;
      if (fileBrowser.visible && editDialog.tool.path) {
        loadFileBrowser();
      }
    };

    const loadFileBrowser = async () => {
      if (!editDialog.tool.path) return;
      
      // 浏览器工具不需要加载文件列表
      if (editDialog.tool.value === 'Browser') {
        fileBrowser.files = [];
        fileBrowser.currentPath = '';
        return;
      }
      
      fileBrowser.loading = true;
      try {
        console.log('前端调用 GetToolDirectory，路径:', editDialog.tool.path);
        const files = await window.go.main.App.GetToolDirectory(editDialog.tool.path);
        fileBrowser.files = files || [];
        fileBrowser.currentPath = editDialog.tool.path;
      } catch (err) {
        console.error('加载文件列表失败:', err);
        ElMessage.error(`加载文件列表失败: ${err}`);
        fileBrowser.files = [];
        
        // 如果目录不存在，可以提供一些帮助信息
        if (err.toString().includes('目录不存在')) {
          ElMessage({
            message: '目录不存在，请检查工具路径是否正确',
            type: 'warning',
            duration: 5000
          });
        }
      } finally {
        fileBrowser.loading = false;
      }
    };

    const refreshFileBrowser = () => {
      loadFileBrowser();
    };

    const onFileClick = (file) => {
      if (file.isDir) {
        // 如果是目录，暂不支持进入子目录
        ElMessage.info('暂不支持浏览子目录');
      } else {
        // 自动填充路径和文件名
        const pathParts = file.path.split('/');
        // 移除文件名，获取目录路径
        pathParts.pop();
        const directoryPath = 'resources/' + pathParts.join('/');
        
        editDialog.tool.path = directoryPath;
        editDialog.tool.fileName = file.name;
        
        // 如果是添加模式且工具名称为空，自动填充工具名称
        if (editDialog.isAddMode && (!editDialog.tool.name || !editDialog.tool.name.trim())) {
          // 去掉文件扩展名作为工具名称
          const nameWithoutExt = file.name.replace(/\.[^/.]+$/, "");
          editDialog.tool.name = nameWithoutExt;
        }
        
        // 根据文件类型自动推荐执行方式
        if (file.name.endsWith('.jar')) {
          if (!editDialog.tool.value || !editDialog.tool.value.startsWith('Java')) {
            editDialog.tool.value = 'Java8'; // 默认推荐Java8
          }
          editDialog.tool.command = '-jar';
        } else if (file.name.endsWith('.py')) {
          if (editDialog.tool.value !== 'custom') {
            editDialog.tool.value = 'custom';
            editDialog.tool.command = 'python {file}';
          }
        } else if (file.name.endsWith('.app') || file.name.endsWith('.exe')) {
          if (editDialog.tool.value !== 'Open' && editDialog.tool.value !== 'custom') {
            editDialog.tool.value = 'Open';
          }
          editDialog.tool.command = '';
        } else {
          // 其他文件类型默认用系统打开
          if (editDialog.tool.value !== 'Open' && editDialog.tool.value !== 'custom') {
            editDialog.tool.value = 'Open';
          }
          editDialog.tool.command = '';
        }
        
        ElMessage.success(`已选择文件: ${file.name}`);
      }
    };

    const getFileIcon = (file) => {
      if (file.isDir) {
        return 'Folder';
      } else if (file.isExecutable) {
        if (file.name.endsWith('.jar')) return 'Coffee';
        if (file.name.endsWith('.app')) return 'Monitor';
        if (file.name.endsWith('.exe')) return 'Platform';
        if (file.name.endsWith('.py')) return 'Document';
        if (file.name.endsWith('.sh')) return 'VideoCamera';
        return 'Files';
      } else {
        return 'Document';
      }
    };

    const getFileIconClass = (file) => {
      if (file.isDir) {
        return 'icon-folder';
      } else if (file.isExecutable) {
        if (file.name.endsWith('.jar')) return 'icon-jar';
        if (file.name.endsWith('.app')) return 'icon-app';
        if (file.name.endsWith('.exe')) return 'icon-exe';
        if (file.name.endsWith('.py')) return 'icon-python';
        if (file.name.endsWith('.sh')) return 'icon-shell';
        return 'icon-executable';
      } else {
        return 'icon-document';
      }
    };

    const getFileTypeBadge = (file) => {
      if (file.name.endsWith('.jar')) return 'JAR';
      if (file.name.endsWith('.app')) return 'APP';
      if (file.name.endsWith('.exe')) return 'EXE';
      if (file.name.endsWith('.py')) return 'PY';
      if (file.name.endsWith('.sh')) return 'SH';
      return 'FILE';
    };

    const isRecommendedFile = (file) => {
      if (file.isDir) return false;
      
      // 推荐主要的可执行文件（通常包含主程序名或版本号）
      const fileName = file.name.toLowerCase();
      const toolName = editDialog.tool?.name?.toLowerCase() || '';
      
      // 如果文件名包含工具名，则推荐
      if (toolName && fileName.includes(toolName.replace(/\s+/g, '-'))) {
        return true;
      }

      // 推荐常见的主程序文件
      const mainPatterns = ['main', 'app', 'client', 'gui', 'tool'];
      return mainPatterns.some(pattern => fileName.includes(pattern));
    };

    const formatFileTime = (timeString) => {
      try {
        const date = new Date(timeString);
        const now = new Date();
        const diffMs = now - date;
        const diffHours = Math.floor(diffMs / (1000 * 60 * 60));
        const diffDays = Math.floor(diffHours / 24);

        if (diffDays > 30) {
          return date.toLocaleDateString('zh-CN', { 
            month: 'short', 
            day: 'numeric' 
          });
        } else if (diffDays > 0) {
          return `${diffDays}天前`;
        } else if (diffHours > 0) {
          return `${diffHours}小时前`;
        } else {
          return '刚刚';
        }
      } catch (e) {
        return '未知';
      }
    };

    // 选择工具目录
    const selectToolDirectory = async () => {
      try {
        const selectedPath = await window.go.main.App.SelectDirectory();
        if (selectedPath) {
          editDialog.tool.path = selectedPath;
          // 重新加载文件浏览器以显示新的目录内容
          loadFileBrowser();
        }
      } catch (err) {
        ElMessage.error(`选择目录失败: ${err}`);
      }
    };

    // 选择工具文件
    const selectToolFile = async () => {
      try {
        const selectedPath = await window.go.main.App.SelectFile();
        if (selectedPath) {
          const pathParts = selectedPath.split('/');
          const fileName = pathParts[pathParts.length - 1];
          const toolPath = pathParts.slice(0, -1).join('/');
          
          editDialog.tool.path = toolPath;
          editDialog.tool.fileName = fileName;
          
          // 如果是添加模式且工具名称为空，自动填充工具名称
          if (editDialog.isAddMode && (!editDialog.tool.name || !editDialog.tool.name.trim())) {
            // 去掉文件扩展名作为工具名称
            const nameWithoutExt = fileName.replace(/\.[^/.]+$/, "");
            editDialog.tool.name = nameWithoutExt;
          }
          
          // 根据文件类型自动推荐执行方式
          if (fileName.endsWith('.jar')) {
            if (!editDialog.tool.value || !editDialog.tool.value.startsWith('Java')) {
              editDialog.tool.value = 'Java8'; // 默认推荐Java8
            }
            editDialog.tool.command = '-jar';
          } else if (fileName.endsWith('.py')) {
            if (editDialog.tool.value !== 'custom') {
              editDialog.tool.value = 'custom';
              editDialog.tool.command = 'python {file}';
            }
          } else if (fileName.endsWith('.app') || fileName.endsWith('.exe')) {
            if (editDialog.tool.value !== 'Open' && editDialog.tool.value !== 'custom') {
              editDialog.tool.value = 'Open';
            }
            editDialog.tool.command = '';
          }
          
          // 重新加载文件浏览器以显示新的目录内容
          if (toolPath) {
            loadFileBrowser();
          }
        }
      } catch (err) {
        ElMessage.error(`选择文件失败: ${err}`);
      }
    };

    const clearFileSelection = () => {
      editDialog.tool.fileName = '';
      editDialog.tool.path = '';
      editDialog.tool.value = '';
      editDialog.tool.command = '';
      ElMessage.info('已清除文件选择');
    };

    // 获取显示路径（从resources开始）
    const getDisplayPath = (fullPath) => {
      if (!fullPath) return '';
      
      // 检查是否为绝对路径
      if (fullPath.startsWith('/') || fullPath.match(/^[A-Za-z]:\\/)) {
        // 绝对路径，显示完整路径但截取过长的部分
        if (fullPath.length > 60) {
          return '...' + fullPath.substring(fullPath.length - 57);
        }
        return fullPath;
      }
      
      // 查找resources在路径中的位置（相对路径处理）
      const resourcesIndex = fullPath.indexOf('resources');
      if (resourcesIndex !== -1) {
        // 从resources开始截取
        return fullPath.substring(resourcesIndex);
      }
      
      // 如果没有找到resources，返回完整路径
      return fullPath;
    };

    // 更新工具路径（处理用户输入）
    const updateToolPath = (value) => {
      // 如果用户输入的是以resources开头的相对路径，我们需要保持原样
      // 如果用户输入的是其他内容，也保持原样
      editDialog.tool.path = value;
    };

    const formatFileSize = (size) => {
      if (size < 1024) return size + ' B';
      if (size < 1024 * 1024) return (size / 1024).toFixed(1) + ' KB';
      if (size < 1024 * 1024 * 1024) return (size / (1024 * 1024)).toFixed(1) + ' MB';
      return (size / (1024 * 1024 * 1024)).toFixed(1) + ' GB';
    };

    const copyToClipboard = async (text) => {
      try {
        await navigator.clipboard.writeText(text);
        ElMessage.success('路径已复制到剪贴板');
      } catch (err) {
        ElMessage.error('复制失败');
      }
    };

    // 侧边栏拖动相关方法
    const startResize = (e) => {
      isResizing.value = true;
      document.addEventListener('mousemove', handleResize, { passive: false });
      document.addEventListener('mouseup', stopResize);
      document.body.style.cursor = 'col-resize';
      document.body.style.userSelect = 'none';
      e.preventDefault();
    };

    const handleResize = (e) => {
      if (!isResizing.value) return;
      
      // 使用 requestAnimationFrame 优化性能
      requestAnimationFrame(() => {
        const newWidth = e.clientX;
        // 扩大可调整范围，增加最小值和最大值的限制
        if (newWidth >= 180 && newWidth <= 500) {
          sidebarWidth.value = newWidth;
        }
      });
      
      e.preventDefault();
    };

    const stopResize = () => {
      isResizing.value = false;
      document.removeEventListener('mousemove', handleResize);
      document.removeEventListener('mouseup', stopResize);
      document.body.style.cursor = '';
      document.body.style.userSelect = '';
      
      // 添加一个小延迟，让过渡效果生效
      setTimeout(() => {
        // 可以在这里添加保存宽度到本地存储的逻辑
        localStorage.setItem('sidebarWidth', sidebarWidth.value.toString());
      }, 100);
    };



    onMounted(async () => {
      // 从本地存储恢复侧边栏宽度
      const savedWidth = localStorage.getItem('sidebarWidth');
      if (savedWidth) {
        const width = parseInt(savedWidth, 10);
        if (width >= 180 && width <= 500) {
          sidebarWidth.value = width;
        }
      }
      
      await loadCategories();
      await loadAllTags();
      
      // 监听命令输出
      window.runtime.EventsOn('command-output', (output) => {
        outputText.value = output;
      });

      // 监听工具添加成功事件
      window.runtime.EventsOn('tool-added', () => {
        loadCategories();
        loadAllTags();
        showAddDialog.value = false;
        ElMessage.success('工具添加成功');
      });

      // 监听工具更新成功事件
      window.runtime.EventsOn('tool-updated', () => {
        if (silentUpdate.value) return;
        loadCategories();
        loadAllTags();
        editDialog.visible = false;
        ElMessage.success('工具修改成功');
      });
      
      // 监听窗口大小变化，确保网格动态响应
      const handleResize = () => {
        // 防抖处理，避免频繁触发
        clearTimeout(handleResize.timer);
        handleResize.timer = setTimeout(() => {
          // 触发CSS重新计算
          const toolsGrid = document.querySelector('.tools-grid');
          if (toolsGrid) {
            // 简单的强制重排
            toolsGrid.style.gridTemplateColumns = toolsGrid.style.gridTemplateColumns;
          }
        }, 100);
      };
      
      window.addEventListener('resize', handleResize);
      
      // 在组件卸载时移除监听器
      onBeforeUnmount(() => {
        window.removeEventListener('resize', handleResize);
        clearTimeout(handleResize.timer);
      });

      // 添加全局点击事件监听器
      document.addEventListener('click', closeContextMenu);
      await loadToolTypes();

      // 添加全局键盘事件监听
      document.addEventListener('keydown', handleGlobalKeydown);
      
      // 添加鼠标移出界面的事件处理
      document.addEventListener('mouseleave', hideTooltip);
    });

    onBeforeUnmount(() => {
      // 移除事件监听器
      document.removeEventListener('click', closeContextMenu);
      
      // 移除全局键盘事件监听
      document.removeEventListener('keydown', handleGlobalKeydown);
      
      // 移除鼠标移出界面的事件处理
      document.removeEventListener('mouseleave', hideTooltip);
    });

    return {
      categories,
      filteredCategories,
      searchQuery,
      showAddDialog,
      activeCategories,
      selectedCategoryName,
      currentTools,
      allTags,
      newTool,
      selectedCategory,
      outputText,
      contextMenu,
      editDialog,
      toolTypes,
      searchInput,

      silentUpdate,

      fileBrowser,
      loadCategories,
      loadAllTags,
      filterTools,
      applySearchFilter,
      executeTool,
      loadToolTypes,
      resetNewToolForm,
      showAddToolDialog,
      scanAndRefreshTools,
      scanCustomDirectory,
      openGitHub,
      selectToolPath,
      showToolMenu,
      closeContextMenu,
      openToolDirectory,
      deleteSelectedTool,
      showEditDialog,
      closeEditDialog,
      submitToolEdit,
      submitToolAdd,
      handleSubmit,
      addTool,
      selectEditToolPath,
      confirmDeleteCategory,
      deleteCategory,
      handleGlobalKeydown,
      executeFirstResult,
      clearSearch,
      handleSpaceSearch,
      onDragEnd,

      selectCategory,
      updateCurrentTools,
      getTotalToolCount,
      getCategoryIcon,
      getToolIcon,
      getWebsiteIcon,
      getTagType,
      formatDate,
      editTool,
      copyToolPath,
      deleteToolConfirm,
      showToolNotes,
      openToolDirectoryByPath,
      searchByTag,
      openToolNote,
      // 笔记编辑
      noteDialog,
      closeNoteDialog,
      togglePreview,
      saveNoteDialog,
      renderMarkdown,
      getToolNotePreview,
      // 分类管理
      showAddCategoryDialog,
      addNewCategory,
      deleteCategoryConfirm,
      startInlineEditCategoryName,
      finishEditCategoryName,
      cancelEditCategoryName,
      editingCategory,
      editingCategoryName,
      updateCategoryName,
      // 图标管理
      iconPopover,
      availableIcons,
      showIconPopover,
      hideIconPopover,
      selectIconFromPopover,
      updateCategoryIconFromPopover,
      // 分类排序
      sortableCategories,
      onCategorySortEnd,
      // Java配置
      javaConfigDialog,
      showJavaConfigDialog,
      selectJavaPath,
      clearJavaPath,
      saveJavaConfig,

      // 侧边栏拖动
      sidebarWidth,
      isResizing,
      startResize,
      // 文件浏览器
      getExecutionTypeLabel,
      isJavaType,
      onExecutionTypeChange,
      toggleFileBrowser,
      loadFileBrowser,
      refreshFileBrowser,
      onFileClick,
      getFileIcon,
      getFileIconClass,
      getFileTypeBadge,
      isRecommendedFile,
      formatFileSize,
      formatFileTime,
      selectToolDirectory,
      selectToolFile,
      clearFileSelection,
      getDisplayPath,
      updateToolPath,
      copyToClipboard,
      // 图标
      Plus,
      Folder,
      Edit,
      Delete,
      CopyDocument,
      Refresh,
      DocumentDelete,
      Select,
      Star,
      CircleCheckFilled,
      Close,
      View,
      Hide
    };
  }
}
</script>

<style>
:root {
  --system-font: -apple-system, BlinkMacSystemFont, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "WenQuanYi Micro Hei", sans-serif;
  --sidebar-width: 250px;
  --primary-color: #409eff;
  --danger-color: #f56c6c;
  --warning-color: #e6a23c;
  --success-color: #67c23a;
}

/* 强制毛玻璃搜索框样式 - 最高优先级 */
.ios-search.el-input .el-input__wrapper,
.ios-search .el-input__wrapper {
  background: rgba(255, 0, 0, 0.3) !important;
  backdrop-filter: blur(20px) saturate(150%) !important;
  -webkit-backdrop-filter: blur(20px) saturate(150%) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  border-radius: 20px !important;
  height: 40px !important;
}

/* 全局重置样式 */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
  overflow: hidden; /* 防止整个页面滚动 */
}

body, 
.app-wrapper,
.sidebar,
.main-content,
.el-button,
.el-input {
  font-family: var(--system-font);
}

.app-wrapper {
  display: flex;
  min-height: 100vh;
  height: 100vh;
  background-color: transparent;
  margin: 0;
  padding: 0;
}

/* 左侧边栏 */
.sidebar {
  width: 280px; /* 默认宽度，将通过样式绑定动态控制 */
  background: rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(10px);
  border-right: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  flex-direction: column;
  position: fixed;
  height: 100vh;
  overflow-y: auto;
  transition: width 0.2s cubic-bezier(0.4, 0.0, 0.2, 1);
}

.sidebar.resizing {
  transition: none; /* 拖拽时禁用过渡效果 */
}

/* 侧边栏拖动分隔线 - 隐形但保留功能 */
.sidebar-resizer {
  position: fixed;
  top: 0;
  width: 6px;
  height: 100vh;
  cursor: col-resize;
  background: transparent;
  z-index: 1000;
  user-select: none;
  margin-left: -3px; /* 居中显示分隔线 */
}

.sidebar-resizer:hover {
  background: transparent; /* 悬停时也保持透明 */
}

.sidebar-resizer:active,
.sidebar-resizer.active {
  background: transparent; /* 拖拽时也保持透明 */
}

/* 拖拽时的全局样式 */
.app-wrapper.resizing {
  user-select: none;
}

.app-wrapper.resizing * {
  cursor: col-resize !important;
}

.sidebar-header {
  padding: 32px 24px 16px 24px; /* 精确的留白比例 */
  display: flex;
  flex-direction: column;
  gap: 16px;
  --wails-draggable: drag; /* 允许拖动窗口 */
}

/* 乔布斯式极简 SpearX 品牌标识 */
.brand-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 16px 0;
  margin-bottom: 16px;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  position: relative;
  width: 100%;
  min-height: 48px;
  overflow: visible;
  background: transparent;
}

/* 简约呼吸光晕 */
.elegant-glow {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 120px;
  height: 120px;
  background: radial-gradient(circle, 
    rgba(255, 255, 255, 0.06) 0%,
    rgba(255, 255, 255, 0.03) 30%,
    rgba(255, 255, 255, 0.01) 60%,
    transparent 80%
  );
  border-radius: 50%;
  animation: gentle-breath 5s ease-in-out infinite;
  pointer-events: none;
  z-index: 1;
}

/* 柔和呼吸动画 */
@keyframes gentle-breath {
  0%, 100% { 
    transform: translate(-50%, -50%) scale(0.7);
    opacity: 0.4;
  }
  50% { 
    transform: translate(-50%, -50%) scale(1.1);
    opacity: 0.8;
  }
}

/* SpearX 极简优雅文字 */
.app-name {
  position: relative;
  z-index: 10;
  font-size: 26px;
  font-weight: 400;
  letter-spacing: 1px;
  font-family: 'SF Pro Display', -apple-system, BlinkMacSystemFont, sans-serif;
  color: rgba(255, 255, 255, 0.92);
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  text-rendering: optimizeLegibility;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  line-height: 1;
  user-select: none;
}

/* 乔布斯式品牌签名 - 极致优雅 */
.brand-signature {
  position: relative;
  z-index: 10;
  font-size: 9px;
  font-weight: 400;
  color: rgba(255, 255, 255, 0.4);
  letter-spacing: 0.3px;
  margin-top: 6px;
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  font-family: 'SF Pro Display', -apple-system, BlinkMacSystemFont, sans-serif;
  text-rendering: optimizeLegibility;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  user-select: none;
  opacity: 0.8;
}

/* 悬停时的精致增强 */
.brand-logo:hover .elegant-glow {
  animation-duration: 3s;
  background: radial-gradient(circle, 
    rgba(255, 255, 255, 0.1) 0%,
    rgba(255, 255, 255, 0.05) 30%,
    rgba(255, 255, 255, 0.02) 60%,
    transparent 80%
  );
  transform: translate(-50%, -50%) scale(1.3);
}

.brand-logo:hover .app-name {
  color: rgba(255, 255, 255, 1);
  letter-spacing: 1.5px;
  transform: translateY(-0.5px);
  text-shadow: 0 0 20px rgba(255, 255, 255, 0.15);
}

.brand-logo:hover .brand-signature {
  color: rgba(255, 255, 255, 0.6);
  letter-spacing: 0.5px;
  opacity: 1;
  transform: translateY(-0.5px);
}

.app-title {
  color: #ffffff;
  font-size: 20px;
  font-weight: 600;
  margin: 0;
  text-align: center;
}

/* 侧边栏图标按钮样式 */
.sidebar-icon-button {
  width: 28px !important;
  height: 28px !important;
  padding: 0 !important;
  border: none !important;
  background: rgba(255, 255, 255, 0.05) !important;
  border-radius: 6px !important;
  color: rgba(255, 255, 255, 0.7) !important;
  font-size: 12px !important;
  transition: all 0.2s ease !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
}

.sidebar-icon-button:hover {
  background: rgba(255, 255, 255, 0.1) !important;
  color: rgba(255, 255, 255, 0.9) !important;
  transform: translateY(-1px) !important;
}

.sidebar-icon-button:active {
  transform: translateY(0) !important;
  background: rgba(255, 255, 255, 0.08) !important;
}

.category-nav {
  flex: 1;
  padding: 0px 24px 24px 24px;
}

.nav-section {
  margin-bottom: 12px;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 10px 16px;
  color: rgba(255, 255, 255, 0.8);
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  font-size: 15px;
  font-weight: 400;
  gap: 10px;
  border-radius: 8px;
  margin: 0 4px 2px 4px;
}

.nav-item:hover {
  background: rgba(255, 255, 255, 0.08);
  color: rgba(255, 255, 255, 0.95);
  transform: translateY(-1px);
}

.nav-item.active {
  background: rgba(255, 255, 255, 0.08);
  color: #ffffff;
  border-radius: 8px;
  margin: 0 8px;
}

/* 分类项样式 */
.category-item {
  padding-right: 35px; /* 为删除按钮预留空间 */
  position: relative;
}

.delete-category-btn {
  position: absolute;
  right: 8px;
  top: 50%;
  transform: translateY(-50%);
  opacity: 0;
  transition: opacity 0.2s ease;
  width: 20px;
  height: 20px;
  min-height: 20px;
  padding: 0;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 4px;
}

.category-item:hover .delete-category-btn {
  opacity: 1;
}

.delete-category-btn:hover {
  background: rgba(255, 0, 0, 0.2);
  border-color: rgba(255, 0, 0, 0.3);
}




.nav-icon {
  font-size: 18px;
  width: 20px;
  text-align: center;
}

.clickable-icon {
  cursor: pointer;
  transition: transform 0.2s ease;
}

.clickable-icon:hover {
  transform: scale(1.1);
}

.nav-text {
  flex: 1;
  font-weight: 500;
}

.editable-text {
  cursor: pointer;
  transition: color 0.2s ease;
}

.editable-text:hover {
  color: rgba(255, 255, 255, 1);
}

.nav-text-input {
  flex: 1;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 4px;
  color: #ffffff;
  padding: 2px 6px;
  font-size: 15px;
  font-weight: 500;
  outline: none;
}

.nav-text-input:focus {
  border-color: #409EFF;
  background: rgba(255, 255, 255, 0.15);
}

/* 图标气泡框样式 */
.icon-popover-content {
  padding: 0;
}

.icon-popover-footer {
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid #ebeef5;
}

.popover-buttons {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  margin-top: 8px;
}

/* 深色主题图标选择器弹出框 */
.dark-icon-popover {
  background: rgba(30, 30, 30, 0.95) !important;
  backdrop-filter: blur(20px);
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.4) !important;
}

.dark-icon-popover .icon-popover-footer {
  border-top: 1px solid rgba(255, 255, 255, 0.1);
}

.dark-icon-popover .el-input :deep(.el-input__wrapper) {
  background: rgba(255, 255, 255, 0.1) !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
  border-radius: 6px;
}

.dark-icon-popover .el-input :deep(.el-input__inner) {
  color: #ffffff !important;
}

.dark-icon-popover .el-input :deep(.el-input__inner)::placeholder {
  color: rgba(255, 255, 255, 0.5) !important;
}

.dark-icon-popover .el-button {
  background: rgba(255, 255, 255, 0.1) !important;
  border: 1px solid rgba(255, 255, 255, 0.2) !important;
  color: #ffffff !important;
}

.dark-icon-popover .el-button:hover {
  background: rgba(255, 255, 255, 0.2) !important;
  border-color: rgba(255, 255, 255, 0.3) !important;
}

.dark-icon-popover .el-button--primary {
  background: rgba(64, 158, 255, 0.8) !important;
  border-color: rgba(64, 158, 255, 0.8) !important;
}

.dark-icon-popover .el-button--primary:hover {
  background: rgba(64, 158, 255, 1) !important;
  border-color: rgba(64, 158, 255, 1) !important;
}

/* 分类拖动排序样式 */
.category-drag-handle {
  color: rgba(255, 255, 255, 0.4);
  cursor: grab;
  padding: 0 4px;
  font-size: 12px;
  transition: color 0.2s ease;
  user-select: none;
}

.category-drag-handle:hover {
  color: rgba(255, 255, 255, 0.8);
}

.category-drag-handle:active {
  cursor: grabbing;
}

.category-ghost {
  opacity: 0.5;
  background: rgba(255, 255, 255, 0.1);
}

.category-chosen {
  background: rgba(255, 255, 255, 0.15);
}

.category-drag {
  transform: rotate(5deg);
}

.nav-count {
  background: rgba(255, 255, 255, 0.2);
  color: #ffffff;
  padding: 2px 6px;
  border-radius: 8px;
  font-size: 10px;
  font-weight: 600;
  min-width: 24px;
  width: 24px;
  text-align: center;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-left: auto;
  flex-shrink: 0;
}



/* 主内容区域 */
.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  transition: margin-left 0.2s cubic-bezier(0.4, 0.0, 0.2, 1);
  height: 100vh;
  overflow: hidden; /* 主容器不滚动，让子元素处理滚动 */
}



.main-content.resizing {
  transition: none; /* 拖拽时禁用过渡效果 */
}

.content-header {
  background: rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding: 16px 20px 16px 20px;
  position: sticky;
  top: 0;
  z-index: 100;
  --wails-draggable: drag; /* 允许拖动窗口 */
}

/* iOS风格搜索栏 - 简洁实用 */
.search-bar {
  display: flex;
  gap: 16px;
  align-items: center;
  width: 100%;
  --wails-draggable: no-drag;
}

.search-wrapper {
  flex: 1;
  max-width: 480px;
}

/* 毛玻璃搜索框 - 最高优先级样式 */
.search-wrapper .el-input.ios-search .el-input__wrapper,
.el-input.ios-search .el-input__wrapper,
.ios-search .el-input__wrapper {
  background: rgba(255, 255, 255, 0.1) !important;
  background-color: rgba(255, 255, 255, 0.1) !important;
  background-image: none !important;
  backdrop-filter: blur(20px) saturate(150%) !important;
  -webkit-backdrop-filter: blur(20px) saturate(150%) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  border-radius: 20px !important;
  box-shadow: 
    0 4px 16px rgba(0, 0, 0, 0.1),
    inset 0 1px 0 rgba(255, 255, 255, 0.15) !important;
  transition: all 0.3s ease !important;
  height: 40px !important;
  position: relative !important;
  overflow: hidden !important;
}

.search-wrapper .el-input.ios-search .el-input__wrapper::before,
.el-input.ios-search .el-input__wrapper::before,
.ios-search .el-input__wrapper::before {
  content: '' !important;
  position: absolute !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  bottom: 0 !important;
  background: linear-gradient(135deg, 
    rgba(255, 255, 255, 0.08) 0%, 
    rgba(255, 255, 255, 0.03) 50%, 
    rgba(255, 255, 255, 0.08) 100%) !important;
  pointer-events: none !important;
  z-index: 1 !important;
}

.search-wrapper .el-input.ios-search .el-input__wrapper:hover,
.el-input.ios-search .el-input__wrapper:hover,
.ios-search .el-input__wrapper:hover {
  background: rgba(255, 255, 255, 0.15) !important;
  background-color: rgba(255, 255, 255, 0.15) !important;
  border-color: rgba(255, 255, 255, 0.25) !important;
  box-shadow: 
    0 8px 24px rgba(0, 0, 0, 0.12),
    inset 0 1px 0 rgba(255, 255, 255, 0.2) !important;
  transform: translateY(-1px) !important;
}

.search-wrapper .el-input.ios-search .el-input__wrapper.is-focus,
.el-input.ios-search .el-input__wrapper.is-focus,
.ios-search .el-input__wrapper.is-focus {
  background: rgba(255, 255, 255, 0.2) !important;
  background-color: rgba(255, 255, 255, 0.2) !important;
  border-color: rgba(10, 132, 255, 0.6) !important;
  box-shadow: 
    0 0 0 2px rgba(10, 132, 255, 0.15),
    0 12px 32px rgba(0, 0, 0, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.25) !important;
  transform: translateY(-1px) !important;
}

.search-wrapper .el-input.ios-search .el-input__inner,
.el-input.ios-search .el-input__inner,
.ios-search .el-input__inner {
  color: rgba(255, 255, 255, 0.9) !important;
  height: 38px !important;
  font-size: 15px !important;
  font-weight: 400 !important;
  font-family: -apple-system, BlinkMacSystemFont, 'SF Pro Display', 'Helvetica Neue', sans-serif !important;
  letter-spacing: -0.2px !important;
  padding-left: 16px !important;
  padding-right: 16px !important;
  border: none !important;
  background: transparent !important;
  text-rendering: optimizeLegibility !important;
  -webkit-font-smoothing: antialiased !important;
  position: relative !important;
  z-index: 2 !important;
}

.search-wrapper .el-input.ios-search .el-input__inner::placeholder,
.el-input.ios-search .el-input__inner::placeholder,
.ios-search .el-input__inner::placeholder {
  color: rgba(255, 255, 255, 0.5) !important;
  font-weight: 300 !important;
}

/* 搜索图标 - macOS风格 */
.ios-search .search-icon {
  color: rgba(255, 255, 255, 0.7);
  font-size: 16px;
  margin-left: 12px;
  margin-right: 6px;
  position: relative;
  z-index: 2;
}

.ios-search:focus-within .search-icon {
  color: rgba(10, 132, 255, 0.9);
}

/* 清除按钮样式 */
.ios-search :deep(.el-input__suffix) {
  right: 12px;
  position: relative;
  z-index: 2;
}

.ios-search :deep(.el-input__clear) {
  color: rgba(255, 255, 255, 0.6);
  font-size: 16px;
}

.ios-search :deep(.el-input__clear):hover {
  color: rgba(255, 255, 255, 0.9);
}

/* 搜索栏按钮样式由 .icon-button 类控制 */

/* 工具容器 */
.tools-container {
  flex: 1;
  padding: 24px;
  overflow-y: auto; /* 允许垂直滚动 */
  height: 0; /* 配合flex:1使用，强制容器计算高度 */
}

.current-category-title {
  font-size: 24px;
  font-weight: 600;
  color: #ffffff;
  margin-bottom: 20px;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
}

/* 工具网格布局 - 动态响应 */
.tools-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 16px;
  width: 100%;
  transition: all 0.3s ease;
}

.tools-grid-inner {
  display: contents;
}

/* 工具卡片 */
.tool-card {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  padding: 10px;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
  display: flex;
  flex-direction: column;
  gap: 6px;
  min-height: 120px;
  cursor: pointer;
  container-type: inline-size; /* 启用容器查询 */
}

.tool-card:hover {
  background: rgba(255, 255, 255, 0.08);
  border-color: rgba(255, 255, 255, 0.2);
  transform: translateY(-4px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.2);
}

/* 工具卡片头部 */
.tool-header {
  display: flex;
  align-items: center;
  gap: 8px;
}

.tool-header .tool-icon {
  width: 28px;
  height: 28px;
  background: rgba(255, 255, 255, 0.1);
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.tool-header .type-icon {
  font-size: 14px;
}

.website-favicon {
  width: 16px;
  height: 16px;
  border-radius: 2px;
}

.tool-title {
  flex: 1;
  font-size: 14px;
  font-weight: 600;
  color: #ffffff;
  line-height: 1.2;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 100%;
}

.tool-footer {
  margin-top: auto;
  opacity: 1;
  transition: all 0.2s ease;
}

.action-group {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  width: 100%;
  padding: 4px;
  overflow: hidden; /* 防止按钮超出卡片 */
}

/* 工具卡片中的按钮样式 - 与顶部按钮一致的简约风格 */
.tool-footer .action-buttons .el-button {
  width: 28px !important;
  height: 28px !important;
  min-width: 24px !important;
  min-height: 24px !important;
  padding: 0 !important;
  background: transparent !important;
  border: none !important;
  border-radius: 6px !important;
  color: rgba(255, 255, 255, 0.6) !important;
  transition: all 0.3s ease !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  flex-shrink: 1 !important; /* 允许按钮适度收缩 */
  flex-grow: 0 !important;
  box-shadow: none !important;
}

.tool-footer .action-buttons .el-button:hover {
  background: rgba(255, 255, 255, 0.08) !important;
  color: rgba(255, 255, 255, 0.9) !important;
  transform: translateY(-2px) !important;
  box-shadow: 
    0 4px 12px rgba(0, 0, 0, 0.3),
    0 2px 6px rgba(255, 255, 255, 0.1) !important;
}

.tool-footer .action-buttons .el-button:active {
  transform: translateY(0) !important;
  background: rgba(255, 255, 255, 0.05) !important;
  box-shadow: 
    0 1px 3px rgba(0, 0, 0, 0.2) !important;
}

.tool-footer .action-buttons .el-button.is-danger {
  background: transparent !important;
  border: none !important;
  color: rgba(245, 108, 108, 0.7) !important;
}

.tool-footer .action-buttons .el-button.is-danger:hover {
  background: rgba(245, 108, 108, 0.1) !important;
  color: #f56c6c !important;
  transform: translateY(-2px) !important;
  box-shadow: 
    0 4px 12px rgba(245, 108, 108, 0.3),
    0 2px 6px rgba(245, 108, 108, 0.1) !important;
}

/* 容器查询 - 基于卡片宽度调整按钮，保持透明风格 */
@container (max-width: 180px) {
  .tool-footer .action-buttons .el-button {
    width: 22px !important;
    height: 22px !important;
    min-width: 18px !important;
    min-height: 18px !important;
    border-radius: 4px !important;
    background: transparent !important;
    border: none !important;
  }
  
  .tool-footer .action-buttons {
    gap: 2px !important;
  }
}

@container (max-width: 200px) {
  .tool-footer .action-buttons .el-button {
    width: 24px !important;
    height: 24px !important;
    min-width: 20px !important;
    min-height: 20px !important;
    border-radius: 5px !important;
    background: transparent !important;
    border: none !important;
  }
  
  .tool-footer .action-buttons {
    gap: 3px !important;
  }
}

/* 主操作按钮容器 */
.action-buttons {
  display: flex;
  align-items: center;
  gap: 8px;
  justify-content: flex-end;
  flex-shrink: 0;
}

/* 工具卡片中的按钮容器特殊样式 - 动态响应 */
.tool-footer .action-buttons {
  gap: 4px !important;
  flex-wrap: nowrap !important;
  overflow: hidden !important;
  justify-content: space-between !important;
  width: 100% !important;
}

/* 简约图标按钮样式 - 仅适用于操作栏和输出区域 */
.action-buttons .icon-button,
.output-header .icon-button {
  width: 32px !important;
  height: 32px !important;
  padding: 0 !important;
  border: none !important;
  background: transparent !important;
  border-radius: 8px !important;
  color: rgba(255, 255, 255, 0.6) !important;
  font-size: 14px !important;
  transition: all 0.3s ease !important;
  display: flex !important;
  align-items: center !important;
  justify-content: center !important;
  box-shadow: none !important;
}

.action-buttons .icon-button:hover,
.output-header .icon-button:hover {
  background: rgba(255, 255, 255, 0.08) !important;
  color: rgba(255, 255, 255, 0.9) !important;
  transform: translateY(-2px) !important;
  box-shadow: 
    0 4px 12px rgba(0, 0, 0, 0.3),
    0 2px 6px rgba(255, 255, 255, 0.1) !important;
}

.action-buttons .icon-button:active,
.output-header .icon-button:active {
  transform: translateY(0) !important;
  background: rgba(255, 255, 255, 0.05) !important;
  box-shadow: 
    0 1px 3px rgba(0, 0, 0, 0.2) !important;
}

/* 工具卡片主体 */
.tool-body {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.tool-path {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12px;
  color: rgba(255, 255, 255, 0.7);
  background: rgba(255, 255, 255, 0.05);
  padding: 6px 8px;
  border-radius: 6px;
}

.path-icon {
  font-size: 14px;
}



.tool-tags {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
}

.clickable-tag {
  cursor: pointer;
  transition: all 0.2s ease;
}

.clickable-tag:hover {
  transform: scale(1.05);
  box-shadow: 0 2px 4px rgba(64, 158, 255, 0.3);
}

.tool-meta {
  display: flex;
  flex-direction: column;
  gap: 4px;
  font-size: 11px;
  color: rgba(255, 255, 255, 0.6);
}

.meta-item {
  display: block;
}



/* 现代化文件浏览器样式 */
.modern-file-browser {
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 12px;
  padding: 0;
  overflow: hidden;
  backdrop-filter: blur(10px);
}

/* 路径显示卡片 */
.path-display-card {
  background: linear-gradient(135deg, rgba(64, 158, 255, 0.1), rgba(103, 194, 58, 0.05));
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  padding: 16px;
}

.path-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 8px;
}

.path-icon {
  font-size: 18px;
  color: #409eff;
}

.path-title {
  font-size: 14px;
  font-weight: 600;
  color: #34495e;
}

.path-content {
  display: flex;
  align-items: center;
  gap: 12px;
}

.path-text {
  flex: 1;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 12px;
  color: #67c23a;
  background: rgba(0, 0, 0, 0.3);
  padding: 8px 12px;
  border-radius: 6px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.copy-btn {
  font-size: 12px;
}

/* 浏览器控制栏 */
.browser-controls {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(0, 0, 0, 0.1);
}

.control-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.file-count {
  font-size: 12px;
  color: #7f8c8d;
  background: rgba(255, 255, 255, 0.08);
  padding: 4px 8px;
  border-radius: 12px;
}

/* 文件浏览器主体 */
.file-browser-main {
  background: rgba(0, 0, 0, 0.05);
}

.browser-content {
  min-height: 100px;
  padding: 12px;
}

/* 空状态 */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 40px 20px;
  color: #95a5a6;
  text-align: center;
}

.empty-icon {
  font-size: 48px;
  margin-bottom: 12px;
  opacity: 0.6;
  color: #bdc3c7;
}

.empty-text {
  font-size: 14px;
  margin: 0;
  color: #7f8c8d;
}

/* 文件列表 */
.files-list {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 6px;
}

/* 文件行 */
.file-row {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 6px 8px;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s ease;
  position: relative;
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid transparent;
}

.file-row:hover {
  background: rgba(255, 255, 255, 0.08);
  border-color: rgba(64, 158, 255, 0.2);
}

.file-row.is-selected {
  background: rgba(64, 158, 255, 0.15);
  border-color: #409eff;
}

.file-row.is-recommended {
  background: rgba(255, 193, 7, 0.08);
  border-color: rgba(255, 193, 7, 0.3);
}

.file-icon-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 20px;
  height: 20px;
  border-radius: 4px;
  background: rgba(255, 255, 255, 0.1);
  flex-shrink: 0;
}

.file-icon {
  font-size: 12px;
  z-index: 1;
}

/* 文件图标颜色 */
.file-icon.icon-jar { color: #ff6b35; }
.file-icon.icon-app { color: #67c23a; }
.file-icon.icon-exe { color: #409eff; }
.file-icon.icon-python { color: #3776ab; }
.file-icon.icon-shell { color: #89e051; }
.file-icon.icon-folder { color: #ffc107; }
.file-icon.icon-executable { color: #e6a23c; }
.file-icon.icon-document { color: rgba(255, 255, 255, 0.6); }

.file-type-badge {
  position: absolute;
  bottom: 0px;
  right: 0px;
  background: #409eff;
  color: white;
  font-size: 6px;
  font-weight: bold;
  padding: 0px 2px;
  border-radius: 2px;
  line-height: 1;
  z-index: 2;
}

.file-info {
  flex: 1;
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 0;
}

.file-name {
  font-size: 11px;
  font-weight: 500;
  color: #2c3e50;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  flex: 1;
}

.file-size {
  font-size: 9px;
  color: #7f8c8d;
  font-weight: 500;
  flex-shrink: 0;
}

.directory-label {
  font-size: 9px;
  color: #67c23a;
  font-weight: 500;
  flex-shrink: 0;
}

.file-actions {
  display: flex;
  align-items: center;
  gap: 4px;
  flex-shrink: 0;
}

.action-icon {
  font-size: 12px;
  opacity: 0.8;
}

.selected-icon {
  color: #67c23a;
}

.recommend-icon {
  color: #ffc107;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 0.6; transform: scale(1); }
  50% { opacity: 1; transform: scale(1.1); }
}





.form-tip {
  font-size: 11px;
  color: #6c757d;
  margin-top: 4px;
  line-height: 1.4;
  background: rgba(255, 193, 7, 0.1);
  padding: 8px 12px;
  border-radius: 6px;
  border-left: 3px solid #ffc107;
}

/* 响应式适配 */
@media (max-width: 768px) {
  .files-list {
    grid-template-columns: 1fr;
    gap: 4px;
  }
  
  .file-row {
    padding: 4px 6px;
    gap: 6px;
  }
  
  .file-icon-wrapper {
    width: 18px;
    height: 18px;
  }
  
  .file-icon {
    font-size: 10px;
  }
  
  .file-name {
    font-size: 10px;
  }
  
  .file-size, .directory-label {
    font-size: 8px;
  }
  
  .action-icon {
    font-size: 10px;
  }
  
  .path-content {
    flex-direction: column;
    gap: 8px;
    align-items: stretch;
  }
}

/* 响应式设计 - 更细致的断点 */
@media (max-width: 1400px) {
  .tools-grid {
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
    gap: 14px;
  }
}

@media (max-width: 1200px) {
  .tools-grid {
    grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
    gap: 12px;
  }
}

@media (max-width: 1000px) {
  .tools-grid {
    grid-template-columns: repeat(auto-fill, minmax(170px, 1fr));
    gap: 12px;
  }
  
  /* 中等屏幕下调整按钮大小，保持透明风格 */
  .tool-footer .action-buttons .el-button {
    width: 26px !important;
    height: 26px !important;
    min-width: 22px !important;
    min-height: 22px !important;
    background: transparent !important;
    border: none !important;
  }
  
  .tool-footer .action-buttons {
    gap: 3px !important;
  }
}

@media (max-width: 768px) {
  .tools-grid {
    grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
    gap: 12px;
  }
  
  .tools-container {
    padding: 16px;
  }
  
  .search-bar {
    flex-wrap: wrap;
    gap: 8px;
  }
  
  .search-bar .el-input {
    max-width: none;
    min-width: 200px;
  }
}

@media (max-width: 600px) {
  .tools-grid {
    grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
    gap: 10px;
  }
  
  /* 小屏幕下进一步缩小按钮，保持透明风格 */
  .tool-footer .action-buttons .el-button {
    width: 24px !important;
    height: 24px !important;
    min-width: 20px !important;
    min-height: 20px !important;
    border-radius: 4px !important;
    background: transparent !important;
    border: none !important;
  }
  
  .tool-footer .action-buttons {
    gap: 2px !important;
  }
}

@media (max-width: 480px) {
  :root {
    --sidebar-width: 100%;
  }
  
  .sidebar {
    position: fixed;
    z-index: 1000;
    transform: translateX(-100%);
    transition: transform 0.3s ease;
  }
  
  .main-content {
    margin-left: 0;
  }
  
  .tools-grid {
    grid-template-columns: 1fr;
    gap: 12px;
  }
  
  .tool-card {
    min-height: 160px;
    padding: 12px;
  }
  
  .tools-container {
    padding: 12px;
  }
}

/* 对话框样式 */
:deep(.el-dialog) {
  margin: 0 !important;
  position: fixed !important;
  max-height: 90vh;
  overflow-y: auto;
}

:deep(.el-dialog__wrapper) {
  overflow: visible !important;
  display: flex !important;
  align-items: center;
  justify-content: center;
}

.el-dialog :deep(.el-form-item__label) {
  color: rgba(255, 255, 255, 0.9);
}

.dialog-footer {
  padding: 10px 20px;
  text-align: right;
}

.dialog-footer .el-button {
  height: 30px;
  font-size: 12px;
  border-radius: 6px;
}

/* 命令输出卡片 */
.output-card {
  margin-top: 24px;
  background-color: rgba(0, 0, 0, 0.9);
  border-radius: 8px;
  color: #ffffff;
}

.output-content {
  white-space: pre-wrap;
  font-family: monospace;
  font-size: 13px;
  color: rgba(255, 255, 255, 0.8);
}

/* 右键菜单 */
.context-menu {
  position: fixed;
  background: rgba(28, 28, 28, 0.95);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 4px 0;
  min-width: 150px;
  z-index: 9999;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.3);
  transition: opacity 0.2s, transform 0.2s;
}

.menu-item {
  padding: 8px 16px;
  display: flex;
  align-items: center;
  gap: 8px;
  color: rgba(255, 255, 255, 0.9);
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 13px;
}

.menu-item:hover {
  background: rgba(255, 255, 255, 0.1);
}

.menu-item.delete {
  color: #ff4d4f;
}

.menu-item.delete:hover {
  background: rgba(255, 77, 79, 0.1);
}

.menu-item .el-icon {
  font-size: 16px;
}

/* Mac风格毛玻璃工具提示 */
.custom-tooltip {
  position: fixed;
  background: linear-gradient(145deg, 
    rgba(255, 255, 255, 0.9), 
    rgba(255, 255, 255, 0.8)
  );
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  color: rgba(0, 0, 0, 0.85);
  padding: 10px 14px;
  border-radius: 8px;
  font-size: 13px;
  font-weight: 500;
  z-index: 9999;
  max-width: 280px;
  word-break: break-word;
  border: 1px solid rgba(255, 255, 255, 0.3);
  box-shadow: 
    0 8px 32px rgba(0, 0, 0, 0.15),
    inset 0 1px 0 rgba(255, 255, 255, 0.4);
  pointer-events: none;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  letter-spacing: 0.2px;
  margin-top: -8px;
}

.custom-tooltip::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(145deg, 
    rgba(255, 255, 255, 0.1), 
    transparent
  );
  border-radius: 7px;
  pointer-events: none;
}

.custom-tooltip::after {
  content: '';
  position: absolute;
  bottom: -6px;
  left: 50%;
  transform: translateX(-50%);
  width: 0;
  height: 0;
  border-left: 6px solid transparent;
  border-right: 6px solid transparent;
  border-top: 6px solid rgba(255, 255, 255, 0.9);
  filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.1));
}

.tooltip-bottom::after {
  bottom: auto;
  top: -6px;
  border-top: none;
  border-bottom: 6px solid rgba(255, 255, 255, 0.9);
}

/* 滚动条隐藏 */
::-webkit-scrollbar {
    width: 0 !important;
}

::-webkit-scrollbar {
  width: 0 !important;
  height: 0;
}

.content-wrapper {
  -ms-overflow-style: none;
  scrollbar-width: none;
  overflow-y: auto;
}

.content-wrapper::-webkit-scrollbar {
  display: none;
}

/* 拖拽样式 */
.ghost {
  opacity: 0.5;
  background: rgba(64, 158, 255, 0.2) !important;
  border: 2px dashed var(--primary-color) !important;
}

/* 优化搜索框的焦点样式 */
.search-bar .el-input :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px var(--primary-color) inset;
  border-color: var(--primary-color);
  background-color: rgba(255, 255, 255, 0.15);
}

/* Java配置对话框样式 */
.java-config-content {
  padding: 16px 0;
}

/* 对话框头部带信息按钮样式 */
.dialog-header-with-info {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  width: 100%;
  padding-right: 40px; /* 为关闭按钮留出空间 */
}

.dialog-title {
  color: rgba(255, 255, 255, 0.95) !important;
  font-weight: 600;
  font-size: 17px;
  letter-spacing: -0.4px;
  line-height: 1.2;
}

.dialog-header-with-info .info-icon {
  margin-left: 8px;
  color: rgba(64, 158, 255, 0.9) !important;
  font-size: 16px !important;
  cursor: help;
  transition: all 0.2s ease;
}

.dialog-header-with-info .info-icon:hover {
  color: rgba(64, 158, 255, 1) !important;
  transform: scale(1.1);
}

/* Java配置提示框样式 */
.java-config-tooltip {
  max-width: 300px !important;
  background: rgba(40, 40, 42, 0.95) !important;
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
  border-radius: 8px !important;
  backdrop-filter: blur(20px) !important;
  -webkit-backdrop-filter: blur(20px) !important;
}

.java-config-tooltip .el-tooltip__content {
  color: rgba(255, 255, 255, 0.9) !important;
  font-size: 13px !important;
  line-height: 1.5 !important;
  padding: 12px !important;
}

/* 优雅的删除确认对话框样式 */
.elegant-confirm-dialog {
  background: rgba(40, 40, 42, 0.95) !important;
  backdrop-filter: blur(30px) saturate(150%) !important;
  -webkit-backdrop-filter: blur(30px) saturate(150%) !important;
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
  border-radius: 16px !important;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.4), 
              0 0 0 1px rgba(255, 255, 255, 0.05) inset !important;
}

.elegant-confirm-dialog .el-message-box__header {
  background: transparent !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08) !important;
  padding: 24px 24px 16px !important;
}

.elegant-confirm-dialog .el-message-box__title {
  color: rgba(255, 255, 255, 0.95) !important;
  font-weight: 600 !important;
  font-size: 18px !important;
  letter-spacing: -0.3px !important;
}

.elegant-confirm-dialog .el-message-box__content {
  background: transparent !important;
  padding: 20px 24px !important;
}

.elegant-confirm-dialog .el-message-box__message {
  color: rgba(255, 255, 255, 0.85) !important;
  font-size: 15px !important;
  line-height: 1.6 !important;
  margin: 0 !important;
}

.elegant-confirm-dialog .el-message-box__btns {
  background: transparent !important;
  border-top: 1px solid rgba(255, 255, 255, 0.08) !important;
  padding: 16px 24px 24px !important;
  text-align: right !important;
}

.elegant-confirm-dialog .el-button {
  border-radius: 10px !important;
  font-weight: 500 !important;
  font-size: 14px !important;
  padding: 10px 20px !important;
  margin-left: 12px !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
}

.elegant-confirm-dialog .el-button--default {
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  color: rgba(255, 255, 255, 0.9) !important;
}

.elegant-confirm-dialog .el-button--default:hover {
  background: rgba(255, 255, 255, 0.15) !important;
  border-color: rgba(255, 255, 255, 0.25) !important;
  transform: translateY(-1px) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3) !important;
}

.elegant-confirm-dialog .el-button--primary {
  background: linear-gradient(135deg, rgba(255, 69, 58, 0.9), rgba(255, 69, 58, 0.8)) !important;
  border: 1px solid rgba(255, 69, 58, 0.8) !important;
  color: #ffffff !important;
}

.elegant-confirm-dialog .el-button--primary:hover {
  background: linear-gradient(135deg, rgba(255, 69, 58, 1), rgba(255, 69, 58, 0.9)) !important;
  border-color: rgba(255, 69, 58, 1) !important;
  transform: translateY(-1px) !important;
  box-shadow: 0 6px 16px rgba(255, 69, 58, 0.4) !important;
}

.elegant-confirm-dialog .el-message-box__close {
  color: rgba(255, 255, 255, 0.6) !important;
  font-size: 16px !important;
  top: 20px !important;
  right: 20px !important;
  width: 32px !important;
  height: 32px !important;
  border-radius: 50% !important;
  background: rgba(255, 255, 255, 0.05) !important;
  transition: all 0.3s ease !important;
}

.elegant-confirm-dialog .el-message-box__close:hover {
  background: rgba(255, 69, 58, 0.8) !important;
  color: #ffffff !important;
  transform: scale(1.1) !important;
}

/* 笔记编辑器样式 */
.note-editor-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.4);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  z-index: 2000;
  display: flex;
  justify-content: flex-end;
  align-items: stretch;
}

.note-editor-card {
  width: 800px;
  height: 100vh;
  background: rgba(40, 40, 42, 0.95);
  backdrop-filter: blur(30px) saturate(150%);
  -webkit-backdrop-filter: blur(30px) saturate(150%);
  border-left: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: -10px 0 30px rgba(0, 0, 0, 0.3);
  display: flex;
  flex-direction: column;
  animation: slideInRight 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes slideInRight {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

.note-card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 24px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.02);
}

.note-card-title {
  display: flex;
  align-items: center;
  gap: 12px;
  color: rgba(255, 255, 255, 0.95);
  font-weight: 600;
  font-size: 18px;
  letter-spacing: -0.3px;
}

.note-card-title .title-icon {
  color: rgba(10, 132, 255, 0.9);
  font-size: 20px;
}

.note-card-actions {
  display: flex;
  align-items: center;
  gap: 8px;
}

.preview-toggle {
  color: rgba(255, 255, 255, 0.7) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  border-radius: 8px !important;
  padding: 8px 16px !important;
  transition: all 0.3s ease !important;
}

.preview-toggle:hover {
  color: rgba(10, 132, 255, 0.9) !important;
  border-color: rgba(10, 132, 255, 0.5) !important;
  background: rgba(10, 132, 255, 0.1) !important;
}

.close-button {
  color: rgba(255, 255, 255, 0.6) !important;
  width: 32px !important;
  height: 32px !important;
  border-radius: 50% !important;
  background: rgba(255, 255, 255, 0.05) !important;
  transition: all 0.3s ease !important;
}

.close-button:hover {
  background: rgba(255, 69, 58, 0.8) !important;
  color: #ffffff !important;
  transform: scale(1.1) !important;
}

.note-card-content {
  flex: 1;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.note-editor-container,
.note-preview-container {
  flex: 1;
  padding: 24px;
  overflow-y: auto;
}

.note-editor-textarea {
  height: 100% !important;
}

.note-editor-textarea .el-textarea__inner {
  background: rgba(255, 255, 255, 0.05) !important;
  border: 1px solid rgba(255, 255, 255, 0.1) !important;
  border-radius: 12px !important;
  color: rgba(255, 255, 255, 0.95) !important;
  font-size: 14px !important;
  line-height: 1.6 !important;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace !important;
  resize: none !important;
  height: 100% !important;
  min-height: 500px !important;
}

.note-editor-textarea .el-textarea__inner:focus {
  border-color: rgba(10, 132, 255, 0.6) !important;
  box-shadow: 0 0 0 2px rgba(10, 132, 255, 0.15) !important;
}

.note-editor-textarea .el-textarea__inner::placeholder {
  color: rgba(255, 255, 255, 0.4) !important;
}

.markdown-preview {
  color: rgba(255, 255, 255, 0.9);
  font-size: 14px;
  line-height: 1.7;
  max-width: none;
}

.markdown-preview h1 {
  color: rgba(255, 255, 255, 0.95);
  font-size: 28px;
  font-weight: 700;
  margin: 24px 0 16px 0;
  border-bottom: 2px solid rgba(10, 132, 255, 0.3);
  padding-bottom: 8px;
}

.markdown-preview h2 {
  color: rgba(255, 255, 255, 0.95);
  font-size: 22px;
  font-weight: 600;
  margin: 20px 0 12px 0;
}

.markdown-preview h3 {
  color: rgba(255, 255, 255, 0.95);
  font-size: 18px;
  font-weight: 600;
  margin: 16px 0 8px 0;
}

.markdown-preview strong {
  color: rgba(255, 255, 255, 0.95);
  font-weight: 600;
}

.markdown-preview em {
  color: rgba(255, 255, 255, 0.8);
  font-style: italic;
}

.markdown-preview code {
  background: rgba(10, 132, 255, 0.15);
  color: rgba(10, 132, 255, 0.9);
  padding: 2px 6px;
  border-radius: 4px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 13px;
}

.markdown-preview pre {
  background: rgba(0, 0, 0, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  padding: 16px;
  overflow-x: auto;
  margin: 12px 0;
}

.markdown-preview pre code {
  background: none;
  color: rgba(255, 255, 255, 0.9);
  padding: 0;
}

.markdown-preview ul {
  margin: 12px 0;
  padding-left: 24px;
}

.markdown-preview li {
  margin: 4px 0;
  color: rgba(255, 255, 255, 0.85);
}

.markdown-preview a {
  color: rgba(10, 132, 255, 0.9);
  text-decoration: none;
  border-bottom: 1px solid rgba(10, 132, 255, 0.3);
  transition: all 0.2s ease;
}

.markdown-preview a:hover {
  color: rgba(10, 132, 255, 1);
  border-bottom-color: rgba(10, 132, 255, 0.6);
}

.note-card-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 20px 24px;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  background: rgba(255, 255, 255, 0.02);
}

.note-info {
  color: rgba(255, 255, 255, 0.6);
  font-size: 13px;
}

.note-tool-info {
  background: rgba(10, 132, 255, 0.15);
  color: rgba(10, 132, 255, 0.9);
  padding: 4px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.note-actions {
  display: flex;
  gap: 12px;
}

.note-actions .el-button {
  border-radius: 8px !important;
  font-weight: 500 !important;
  padding: 10px 20px !important;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1) !important;
}

.note-actions .el-button:not(.el-button--primary) {
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  color: rgba(255, 255, 255, 0.9) !important;
}

.note-actions .el-button:not(.el-button--primary):hover {
  background: rgba(255, 255, 255, 0.15) !important;
  border-color: rgba(255, 255, 255, 0.25) !important;
  transform: translateY(-1px) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3) !important;
}

.note-actions .el-button--primary {
  background: rgba(10, 132, 255, 0.9) !important;
  border: 1px solid rgba(10, 132, 255, 0.9) !important;
  color: #ffffff !important;
}

.note-actions .el-button--primary:hover {
  background: rgba(10, 132, 255, 1) !important;
  border-color: rgba(10, 132, 255, 1) !important;
  transform: translateY(-1px) !important;
  box-shadow: 0 6px 16px rgba(10, 132, 255, 0.4) !important;
}

.config-description {
  margin-bottom: 24px;
  animation: slideDown 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.description-content {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 16px;
  background: rgba(64, 158, 255, 0.15);
  border: 1px solid rgba(64, 158, 255, 0.3);
  border-radius: 8px;
  font-size: 14px;
  color: #ffffff !important;
  font-weight: 500;
}

.description-content .info-icon {
  color: var(--primary-color);
  font-size: 16px;
  flex-shrink: 0;
}

/* Java配置表单 - 暗色系样式优化 */
.java-config-form .el-form-item {
  margin-bottom: 24px;
}

.java-config-form .el-form-item__label {
  color: rgba(255, 255, 255, 0.9) !important;
  font-weight: 500;
  font-size: 14px;
  letter-spacing: -0.2px;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}

.java-config-form .el-input {
  border-radius: 8px;
}

.java-config-form .el-input__wrapper {
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.12) !important;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  box-shadow: none;
}

.java-config-form .el-input__wrapper:hover {
  background: rgba(255, 255, 255, 0.12) !important;
  border-color: rgba(255, 255, 255, 0.2) !important;
}

.java-config-form .el-input__wrapper.is-focus {
  background: rgba(255, 255, 255, 0.15) !important;
  border-color: rgba(10, 132, 255, 0.8) !important;
  box-shadow: 0 0 0 2px rgba(10, 132, 255, 0.15);
}

.java-config-form .el-input__inner {
  color: rgba(255, 255, 255, 0.95) !important;
  font-size: 14px;
  font-weight: 400;
  background: transparent !important;
}

.java-config-form .el-input__inner::placeholder {
  color: rgba(255, 255, 255, 0.5) !important;
}

.java-config-form .el-button {
  border-radius: 8px;
  font-weight: 500;
  font-size: 14px;
  letter-spacing: -0.2px;
  padding: 8px 16px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.12) !important;
  color: rgba(255, 255, 255, 0.9) !important;
  box-shadow: none;
}

.java-config-form .el-button:hover {
  background: rgba(255, 255, 255, 0.15) !important;
  border-color: rgba(255, 255, 255, 0.25) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  transform: translateY(-1px);
}

/* ========================================
   Mac风格毛玻璃对话框统一样式系统
   ======================================== */

/* 暗色系简洁对话框样式 */
.glassmorphism-dialog {
  background: rgba(40, 40, 42, 0.95) !important;
  backdrop-filter: blur(30px) saturate(150%);
  -webkit-backdrop-filter: blur(30px) saturate(150%);
  border: none;
  border-radius: 12px;
  box-shadow: none;
  overflow: hidden;
  animation: dialogFadeIn 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes dialogFadeIn {
  from {
    opacity: 0;
    transform: scale(0.95) translateY(-10px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
  }
}

/* 简洁对话框标题区域 */
.glassmorphism-dialog .el-dialog__header {
  background: transparent !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 12px 12px 0 0;
  padding: 20px 24px 16px;
  position: relative;
}

.glassmorphism-dialog .el-dialog__header::before {
  display: none;
}

.glassmorphism-dialog .el-dialog__title {
  color: rgba(255, 255, 255, 0.95) !important;
  font-weight: 600;
  font-size: 17px;
  letter-spacing: -0.4px;
  line-height: 1.2;
}

/* 暗色系对话框内容区域 */
.glassmorphism-dialog .el-dialog__body {
  background: transparent !important;
  color: rgba(255, 255, 255, 0.9) !important;
  padding: 24px;
  position: relative;
}

/* 简洁对话框底部区域 */
.glassmorphism-dialog .el-dialog__footer {
  background: transparent !important;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 0 0 12px 12px;
  padding: 16px 24px 20px;
  position: relative;
}

.glassmorphism-dialog .el-dialog__footer::before {
  display: none;
}

/* 关闭按钮 - 悬停阴影效果 */
.glassmorphism-dialog .el-dialog__headerbtn {
  top: 16px;
  right: 20px;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.08);
  border: none;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: none;
}

.glassmorphism-dialog .el-dialog__headerbtn:hover {
  background: rgba(255, 69, 58, 0.8);
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(255, 69, 58, 0.3);
}

.glassmorphism-dialog .el-dialog__close {
  color: rgba(255, 255, 255, 0.8) !important;
  font-size: 14px;
  font-weight: 500;
  transition: color 0.2s ease;
}

.glassmorphism-dialog .el-dialog__headerbtn:hover .el-dialog__close {
  color: #ffffff !important;
}

/* Java配置对话框 - 简洁样式 */
.el-dialog.java-config-dialog {
  background: rgba(40, 40, 42, 0.95) !important;
  backdrop-filter: blur(30px) saturate(150%);
  -webkit-backdrop-filter: blur(30px) saturate(150%);
  border: none;
  border-radius: 12px;
  box-shadow: none;
  overflow: hidden;
  animation: dialogFadeIn 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.el-dialog.java-config-dialog .el-dialog__header {
  background: transparent !important;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 12px 12px 0 0;
  padding: 20px 24px 16px;
  position: relative;
}

.el-dialog.java-config-dialog .el-dialog__header::before {
  display: none;
}

.el-dialog.java-config-dialog .el-dialog__title {
  color: rgba(255, 255, 255, 0.95) !important;
  font-weight: 600;
  font-size: 17px;
  letter-spacing: -0.4px;
  line-height: 1.2;
}

.el-dialog.java-config-dialog .el-dialog__body {
  background: transparent !important;
  color: rgba(255, 255, 255, 0.9) !important;
  padding: 24px;
  position: relative;
}

.el-dialog.java-config-dialog .el-dialog__footer {
  background: transparent !important;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 0 0 12px 12px;
  padding: 16px 24px 20px;
  position: relative;
}

.el-dialog.java-config-dialog .el-dialog__footer::before {
  display: none;
}

.el-dialog.java-config-dialog .el-dialog__headerbtn {
  top: 16px;
  right: 20px;
  width: 28px;
  height: 28px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.08);
  border: none;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: none;
}

.el-dialog.java-config-dialog .el-dialog__headerbtn:hover {
  background: rgba(255, 69, 58, 0.8);
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(255, 69, 58, 0.3);
}

.el-dialog.java-config-dialog .el-dialog__close {
  color: rgba(255, 255, 255, 0.8) !important;
  font-size: 14px;
  font-weight: 500;
  transition: color 0.2s ease;
}

.el-dialog.java-config-dialog .el-dialog__headerbtn:hover .el-dialog__close {
  color: #ffffff !important;
}

/* Java配置对话框底部按钮 - 悬停阴影效果 */
.el-dialog.java-config-dialog .dialog-footer .el-button {
  border-radius: 8px;
  font-weight: 500;
  font-size: 14px;
  letter-spacing: -0.2px;
  padding: 10px 20px;
  min-width: 80px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  box-shadow: none;
}

.el-dialog.java-config-dialog .dialog-footer .el-button:not(.el-button--primary) {
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.12) !important;
  color: rgba(255, 255, 255, 0.9) !important;
}

.el-dialog.java-config-dialog .dialog-footer .el-button:not(.el-button--primary):hover {
  background: rgba(255, 255, 255, 0.15) !important;
  border-color: rgba(255, 255, 255, 0.25) !important;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  transform: translateY(-1px);
}

.el-dialog.java-config-dialog .dialog-footer .el-button--primary {
  background: rgba(10, 132, 255, 0.9) !important;
  border: 1px solid rgba(10, 132, 255, 0.9) !important;
  color: #ffffff !important;
}

.el-dialog.java-config-dialog .dialog-footer .el-button--primary:hover {
  background: rgba(10, 132, 255, 1) !important;
  border-color: rgba(10, 132, 255, 1) !important;
  box-shadow: 0 6px 16px rgba(10, 132, 255, 0.4);
  transform: translateY(-1px);
}

/* Java路径输入组样式 */
.java-path-input-group {
  display: flex;
  align-items: center;
  gap: 12px;
  width: 100%;
}

.java-path-input {
  flex: 1;
  min-width: 0; /* 允许收缩 */
}

.java-clear-button {
  flex-shrink: 0; /* 防止按钮收缩 */
  min-width: 40px !important;
  padding: 6px 8px !important;
}

/* 右侧编辑工具卡片样式 */
/* 右侧编辑工具卡片 - 暗色系Mac风格毛玻璃 */
.edit-tool-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(30px) saturate(180%);
  -webkit-backdrop-filter: blur(30px) saturate(180%);
  z-index: 2000;
  display: flex;
  justify-content: flex-end;
  align-items: stretch;
  animation: overlayFadeIn 0.25s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes overlayFadeIn {
  from {
    opacity: 0;
    backdrop-filter: blur(0px);
    -webkit-backdrop-filter: blur(0px);
  }
  to {
    opacity: 1;
    backdrop-filter: blur(30px) saturate(180%);
    -webkit-backdrop-filter: blur(30px) saturate(180%);
  }
}

.edit-tool-card {
  width: 520px;
  background: rgba(28, 28, 30, 0.95);
  backdrop-filter: blur(50px) saturate(180%) brightness(1.1);
  -webkit-backdrop-filter: blur(50px) saturate(180%) brightness(1.1);
  border: 0.5px solid rgba(255, 255, 255, 0.15);
  border-right: none;
  box-shadow: 
    -25px 0 50px -12px rgba(0, 0, 0, 0.8),
    0 0 0 0.5px rgba(255, 255, 255, 0.08) inset,
    0 1px 0 0 rgba(255, 255, 255, 0.1) inset,
    0 0 30px rgba(0, 0, 0, 0.5);
  display: flex;
  flex-direction: column;
  height: 100vh;
  overflow: hidden;
  position: relative;
  border-radius: 16px 0 0 16px;
  animation: cardSlideIn 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

@keyframes cardSlideIn {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

.edit-tool-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, 
    rgba(255, 255, 255, 0.06) 0%,
    rgba(255, 255, 255, 0.02) 50%,
    transparent 100%
  );
  pointer-events: none;
  z-index: 0;
  border-radius: 16px 0 0 16px;
}

@keyframes slideInRight {
  from {
    transform: translateX(100%);
    opacity: 0;
  }
  to {
    transform: translateX(0);
  opacity: 1;
  }
}

/* 编辑卡片标题 - 暗色系样式 */
.edit-card-header {
  padding: 20px 24px 16px;
  background: linear-gradient(135deg, 
    rgba(255, 255, 255, 0.08) 0%,
    rgba(255, 255, 255, 0.03) 100%);
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  border-bottom: 0.5px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px 0 0 0;
  display: flex;
  align-items: center;
  justify-content: space-between;
  position: relative;
}

.edit-card-header::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, 
    transparent,
    rgba(255, 255, 255, 0.2) 50%,
    transparent);
  pointer-events: none;
}

.edit-card-title {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 17px;
  font-weight: 600;
  color: rgba(255, 255, 255, 0.95);
  letter-spacing: -0.4px;
  line-height: 1.2;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
  position: relative;
  z-index: 1;
}

/* 标题图标 - 暗色系样式 */
.title-icon {
  font-size: 20px;
  color: rgba(10, 132, 255, 0.9);
  filter: drop-shadow(0 1px 3px rgba(10, 132, 255, 0.3));
}

/* 统一的关闭按钮 - 悬停阴影效果 */
.close-btn {
  color: rgba(255, 255, 255, 0.8) !important;
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.12) !important;
  border-radius: 50%;
  width: 28px;
  height: 28px;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  font-size: 14px;
  font-weight: 500;
  position: relative;
  z-index: 1;
  box-shadow: none;
}

.close-btn:hover {
  background: rgba(255, 69, 58, 0.9) !important;
  border-color: rgba(255, 69, 58, 0.9) !important;
  color: #ffffff !important;
  transform: scale(1.05);
  box-shadow: 0 4px 12px rgba(255, 69, 58, 0.3);
}

/* 编辑卡片表单 - 暗色系样式 */
.edit-tool-card .el-form {
  padding: 24px;
  flex: 1;
  overflow-y: auto;
}

.edit-tool-card .el-form-item {
  margin-bottom: 24px;
}

.edit-tool-card .el-form-item__label {
  color: rgba(255, 255, 255, 0.9) !important;
  font-weight: 500;
  font-size: 14px;
  letter-spacing: -0.2px;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}

.edit-tool-card .el-input :deep(.el-input__wrapper) {
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  border-radius: 8px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.2);
}

.edit-tool-card .el-input :deep(.el-input__wrapper:hover) {
  background: rgba(255, 255, 255, 0.12) !important;
  border-color: rgba(255, 255, 255, 0.25) !important;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.3);
}

.edit-tool-card .el-input :deep(.el-input__wrapper.is-focus) {
  background: rgba(255, 255, 255, 0.15) !important;
  border-color: rgba(10, 132, 255, 0.8) !important;
  box-shadow: 
    0 0 0 3px rgba(10, 132, 255, 0.2),
    0 2px 12px rgba(0, 0, 0, 0.4);
}

.edit-tool-card .el-input :deep(.el-input__inner) {
  color: rgba(255, 255, 255, 0.95) !important;
  font-size: 14px;
  font-weight: 400;
}

.edit-tool-card .el-input :deep(.el-input__inner::placeholder) {
  color: rgba(255, 255, 255, 0.5) !important;
}

.edit-tool-card .el-select :deep(.el-input__wrapper) {
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  border-radius: 8px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
}

.edit-tool-card .el-textarea :deep(.el-textarea__inner) {
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.15) !important;
  border-radius: 8px;
  color: rgba(255, 255, 255, 0.95) !important;
  font-size: 14px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
}

.edit-tool-card .el-textarea :deep(.el-textarea__inner:focus) {
  background: rgba(255, 255, 255, 0.15) !important;
  border-color: rgba(10, 132, 255, 0.8) !important;
  box-shadow: 
    0 0 0 3px rgba(10, 132, 255, 0.2),
    0 2px 12px rgba(0, 0, 0, 0.4);
}

/* 执行选项卡片样式 */
.execution-option {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 4px 0;
}

.option-icon {
  font-size: 16px;
  flex-shrink: 0;
}

.java-icon {
  color: #f89820;
}

.open-icon {
  color: #67c23a;
}

.terminal-icon {
  color: #e6a23c;
}

.custom-icon {
  color: #909399;
}

.browser-icon {
  color: #409eff;
}

.option-content {
  flex: 1;
  min-width: 0;
}

.option-title {
  font-size: 14px;
  font-weight: 500;
  color: #303133;
  line-height: 1.2;
}

.option-desc {
  font-size: 12px;
  color: #909399;
  line-height: 1.2;
  margin-top: 2px;
}

/* 修改文件浏览器在卡片中的样式 */
.edit-tool-card .modern-file-browser {
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 2px;
  overflow: hidden;
}

.edit-tool-card .browser-controls {
  background: rgba(255, 255, 255, 0.05);
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  padding: 12px 16px;
}

.edit-tool-card .files-list {
  max-height: 300px;
  overflow-y: auto;
  padding: 8px;
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.edit-tool-card .file-row {
  background: rgba(255, 255, 255, 0.02);
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: 2px;
  transition: all 0.2s ease;
  padding: 8px 12px;
  cursor: pointer;
  width: 100%;
}

.edit-tool-card .file-row:hover {
  background: rgba(255, 255, 255, 0.08);
  border-color: rgba(64, 158, 255, 0.3);
  transform: translateY(-1px);
}

.edit-tool-card .file-row.is-selected {
  background: rgba(64, 158, 255, 0.15);
  border-color: #409eff;
  box-shadow: 0 0 10px rgba(64, 158, 255, 0.2);
}

.edit-tool-card .file-name {
  color: rgba(255, 255, 255, 0.9) !important;
  font-weight: 500;
  font-size: 13px;
}

.edit-tool-card .file-meta {
  color: rgba(255, 255, 255, 0.5) !important;
  font-size: 11px;
}

/* 编辑卡片底部按钮 - 悬停阴影效果 */
.edit-card-footer {
  padding: 16px 24px 20px;
  background: transparent;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 0 0 0 16px;
  display: flex;
  gap: 12px;
  justify-content: flex-end;
  position: relative;
}

.edit-card-footer::before {
  display: none;
}

.edit-card-footer .el-button {
  min-width: 100px;
  border-radius: 8px;
  font-weight: 500;
  font-size: 14px;
  letter-spacing: -0.2px;
  padding: 10px 20px;
  position: relative;
  z-index: 1;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
  box-shadow: none;
}

.edit-card-footer .el-button:not(.el-button--primary) {
  background: rgba(255, 255, 255, 0.08) !important;
  border: 1px solid rgba(255, 255, 255, 0.12) !important;
  color: rgba(255, 255, 255, 0.9) !important;
}

.edit-card-footer .el-button:not(.el-button--primary):hover {
  background: rgba(255, 255, 255, 0.15) !important;
  border-color: rgba(255, 255, 255, 0.25) !important;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.edit-card-footer .el-button--primary {
  background: rgba(10, 132, 255, 0.9) !important;
  border: 1px solid rgba(10, 132, 255, 0.9) !important;
  color: #ffffff !important;
}

.edit-card-footer .el-button--primary:hover {
  background: rgba(10, 132, 255, 1) !important;
  border-color: rgba(10, 132, 255, 1) !important;
  transform: translateY(-1px);
  box-shadow: 0 6px 16px rgba(10, 132, 255, 0.4);
}

@media (max-width: 768px) {
  .edit-tool-card {
    width: 100%;
    animation: slideInUp 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  }
  
  @keyframes slideInUp {
    from {
      transform: translateY(100%);
      opacity: 0;
    }
    to {
      transform: translateY(0);
      opacity: 1;
    }
  }
  
  .edit-card-footer {
    flex-direction: column-reverse;
  }
  
  .edit-card-footer .el-button {
    width: 100%;
  }
}








/* 紧凑型图标选择器样式 */
.compact-icon-selector {
  padding: 0;
}

/* ========================================
   Mac原生风格对话框优化
   ======================================== */

/* Mac原生风格设置对话框样式 */
.mac-native-dialog {
  background: rgba(50, 50, 52, 0.95) !important;
  backdrop-filter: blur(60px) saturate(200%) brightness(1.2);
  -webkit-backdrop-filter: blur(60px) saturate(200%) brightness(1.2);
  border: 0.5px solid rgba(255, 255, 255, 0.12);
  border-radius: 12px;
  box-shadow: 
    0 30px 60px -12px rgba(0, 0, 0, 0.8),
    0 0 0 0.5px rgba(255, 255, 255, 0.06) inset,
    0 1px 0 0 rgba(255, 255, 255, 0.08) inset;
  overflow: hidden;
  animation: macDialogFadeIn 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

@keyframes macDialogFadeIn {
  from {
    opacity: 0;
    transform: scale(0.92) translateY(-20px);
    filter: blur(8px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
    filter: blur(0px);
  }
}

/* Mac风格选择器样式 */
.mac-select {
  background: rgba(70, 70, 72, 0.8) !important;
  border: 1px solid rgba(255, 255, 255, 0.08) !important;
  border-radius: 6px;
  color: rgba(255, 255, 255, 0.95) !important;
  font-size: 14px;
  font-weight: 500;
  padding: 8px 32px 8px 12px;
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
  position: relative;
}

.mac-select:hover {
  background: rgba(70, 70, 72, 0.9) !important;
  border-color: rgba(255, 255, 255, 0.15) !important;
}

.mac-select:focus {
  border-color: rgba(10, 132, 255, 0.8) !important;
  box-shadow: 0 0 0 3px rgba(10, 132, 255, 0.15);
}

/* Mac风格下拉箭头 */
.mac-select::after {
  content: '';
  position: absolute;
  right: 12px;
  top: 50%;
  transform: translateY(-50%);
  width: 0;
  height: 0;
  border-left: 4px solid transparent;
  border-right: 4px solid transparent;
  border-top: 5px solid rgba(255, 255, 255, 0.6);
  pointer-events: none;
}

/* Mac风格复选框 */
.mac-checkbox {
  appearance: none;
  width: 18px;
  height: 18px;
  border: 1.5px solid rgba(255, 255, 255, 0.3);
  border-radius: 3px;
  background: rgba(70, 70, 72, 0.6);
  position: relative;
  cursor: pointer;
  transition: all 0.2s ease;
}

.mac-checkbox:checked {
  background: linear-gradient(135deg, #007AFF, #0056CC);
  border-color: #007AFF;
}

.mac-checkbox:checked::after {
  content: '';
  position: absolute;
  left: 3px;
  top: 0px;
  width: 4px;
  height: 8px;
  border: solid white;
  border-width: 0 2px 2px 0;
  transform: rotate(45deg);
}

/* Mac风格滑块 */
.mac-slider {
  -webkit-appearance: none;
  appearance: none;
  width: 100%;
  height: 4px;
  border-radius: 2px;
  background: rgba(70, 70, 72, 0.8);
  outline: none;
  margin: 20px 0;
}

.mac-slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: linear-gradient(135deg, #ffffff, #f0f0f0);
  cursor: pointer;
  border: 2px solid rgba(0, 0, 0, 0.1);
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.3);
}

.mac-slider::-webkit-slider-thumb:hover {
  transform: scale(1.1);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.4);
}

/* Mac风格按钮 */
.mac-button {
  background: rgba(70, 70, 72, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 6px;
  color: rgba(255, 255, 255, 0.9);
  font-size: 13px;
  font-weight: 500;
  padding: 6px 16px;
  cursor: pointer;
  transition: all 0.2s ease;
  backdrop-filter: blur(20px);
  -webkit-backdrop-filter: blur(20px);
}

.mac-button:hover {
  background: rgba(70, 70, 72, 0.95);
  border-color: rgba(255, 255, 255, 0.2);
  transform: translateY(-1px);
}

.mac-button:active {
  transform: translateY(0);
  background: rgba(60, 60, 62, 0.9);
}

/* Mac风格分组标题 */
.mac-section-title {
  color: rgba(255, 255, 255, 0.9);
  font-size: 13px;
  font-weight: 600;
  margin: 24px 0 12px 0;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}

/* Mac风格描述文字 */
.mac-description {
  color: rgba(255, 255, 255, 0.7);
  font-size: 12px;
  line-height: 1.4;
  margin-top: 4px;
}

/* ========================================
   通用的毛玻璃对话框优化
   ======================================== */

/* 优化所有对话框的文字排版 - 暗色系 */
:deep(.el-dialog .el-form-item__label) {
  color: rgba(255, 255, 255, 0.9) !important;
  font-weight: 500;
  font-size: 14px;
  letter-spacing: -0.2px;
  line-height: 1.4;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.3);
}

/* 优化描述文字 - 简洁样式 */
.config-description {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  padding: 16px;
  margin-bottom: 24px;
}

.config-description .info-icon {
  color: rgba(10, 132, 255, 0.9);
  margin-right: 8px;
  font-size: 16px;
  filter: drop-shadow(0 1px 2px rgba(10, 132, 255, 0.3));
}

.config-description span {
  color: rgba(255, 255, 255, 0.85);
  font-size: 14px;
  line-height: 1.5;
  letter-spacing: -0.1px;
}

/* 优化按钮的统一样式 */
:deep(.el-button) {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  letter-spacing: -0.2px;
  font-weight: 500;
}

/* 增强暗色系毛玻璃效果的视觉层次 */
.glassmorphism-dialog::after {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(
    circle at 20% 20%,
    rgba(255, 255, 255, 0.03) 0%,
    transparent 60%
  );
  pointer-events: none;
  z-index: 0;
}

/* 动画优化 */
@keyframes dialogEnter {
  from {
    opacity: 0;
    transform: scale(0.9) translateY(-20px);
    filter: blur(10px);
  }
  to {
    opacity: 1;
    transform: scale(1) translateY(0);
    filter: blur(0px);
  }
}

.glassmorphism-dialog {
  animation: dialogEnter 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.icon-grid-compact {
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  gap: 6px;
  max-height: 280px;
  overflow-y: auto;
  padding: 12px;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.icon-item-compact {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  font-size: 18px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.15s ease;
  user-select: none;
}

.icon-item-compact:hover {
  background: rgba(255, 255, 255, 0.12);
  border-color: rgba(64, 158, 255, 0.4);
  transform: scale(1.1);
}

.icon-item-compact.active {
  background: rgba(64, 158, 255, 0.25);
  border-color: #409eff;
  box-shadow: 0 0 8px rgba(64, 158, 255, 0.4);
  transform: scale(1.1);
}

/* 滚动条样式 */
.icon-grid-compact::-webkit-scrollbar {
  width: 4px;
}

.icon-grid-compact::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.05);
  border-radius: 2px;
}

.icon-grid-compact::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 2px;
}

.icon-grid-compact::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.3);
}



/* 代码块样式 */
.code-block-container {
  position: relative;
  margin: 8px 0;
}

.code-copy-btn {
  position: absolute;
  top: 8px;
  right: 8px;
  z-index: 10;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 4px 8px;
  min-width: 36px;
  height: 22px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(59, 130, 246, 0.6);
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.2s ease;
  color: rgba(255, 255, 255, 0.9);
  font-size: 10px;
  font-weight: 500;
  white-space: nowrap;
  backdrop-filter: blur(4px);
}

.code-copy-btn:hover {
  background: rgba(59, 130, 246, 0.2);
  border-color: rgba(59, 130, 246, 0.8);
  color: rgba(255, 255, 255, 1);
  transform: scale(1.02);
}

.code-copy-btn:active {
  transform: scale(0.98);
  background: rgba(255, 255, 255, 0.1);
}

.code-block-container pre {
  margin: 0;
  padding: 32px 50px 12px 16px; /* 右上角留出空间给复制按钮 */
  background: rgba(0, 0, 0, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 8px;
  overflow-x: auto;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 13px;
  line-height: 1.4;
}

.code-block-container code {
  background: transparent;
  color: rgba(255, 255, 255, 0.9);
  font-family: inherit;
  padding: 0;
  border-radius: 0;
  white-space: pre;
}

/* 语法高亮样式 */
.code-block-container .language-shell,
.code-block-container .language-bash {
  color: #a6e3a1; /* 绿色 */
}

.code-block-container .language-python {
  color: #f9e2af; /* 黄色 */
}

.code-block-container .language-javascript,
.code-block-container .language-js {
  color: #fab387; /* 橙色 */
}

.code-block-container .language-java {
  color: #f38ba8; /* 粉色 */
}

.code-block-container .language-go {
  color: #89dceb; /* 青色 */
}

.code-block-container .language-sql {
  color: #cba6f7; /* 紫色 */
}

/* 行内代码样式 */
.inline-code {
  background: rgba(255, 255, 255, 0.1);
  color: rgba(255, 255, 255, 0.9);
  padding: 2px 6px;
  border-radius: 4px;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-size: 0.9em;
}

/* Markdown预览区域样式增强 */
.markdown-preview {
  color: rgba(255, 255, 255, 0.9);
  line-height: 1.6;
}

.markdown-preview h1,
.markdown-preview h2,
.markdown-preview h3 {
  color: rgba(255, 255, 255, 0.95);
  margin: 16px 0 8px 0;
}

.markdown-preview h1 {
  font-size: 1.8em;
  border-bottom: 2px solid rgba(255, 255, 255, 0.1);
  padding-bottom: 8px;
}

.markdown-preview h2 {
  font-size: 1.5em;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  padding-bottom: 4px;
}

.markdown-preview h3 {
  font-size: 1.3em;
}

.markdown-preview ul {
  margin: 8px 0;
  padding-left: 20px;
}

.markdown-preview li {
  margin: 4px 0;
  color: rgba(255, 255, 255, 0.85);
}

.markdown-preview a {
  color: #409eff;
  text-decoration: none;
}

.markdown-preview a:hover {
  text-decoration: underline;
}

.markdown-preview strong {
  color: rgba(255, 255, 255, 0.95);
  font-weight: 600;
}

.markdown-preview em {
  color: rgba(255, 255, 255, 0.9);
  font-style: italic;
}

/* 开发者信息样式 */
.developer-info {
  position: fixed;
  bottom: 20px;
  right: 20px;
  z-index: 9999;
  background: rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(15px);
  padding: 10px 16px;
  border-radius: 24px;
  border: 1px solid rgba(255, 255, 255, 0.08);
  transition: all 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94);
  opacity: 0.6;
  cursor: pointer;
}

.developer-info:hover {
  opacity: 0.9;
  background: rgba(0, 0, 0, 0.7);
  transform: translateY(-1px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.2);
  border-color: rgba(255, 255, 255, 0.12);
}

.dev-text {
  color: rgba(255, 255, 255, 0.9);
  font-size: 11px;
  font-weight: 500;
  white-space: nowrap;
}

.dev-name {
  color: #00d4ff;
  font-weight: 700;
  text-shadow: 0 0 8px rgba(0, 212, 255, 0.5);
  transition: all 0.3s ease;
}

.dev-name:hover {
  text-shadow: 0 0 12px rgba(0, 212, 255, 0.8);
}

</style>


