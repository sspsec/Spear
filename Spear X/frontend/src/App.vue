
<template>
  <div class="app-wrapper">
    <div class="container" @keyup.space="handleSpaceSearch" tabindex="0">
      <div class="header">
        <div class="search-bar">
          <el-input
            ref="searchInput"
            v-model="searchQuery"
            placeholder="搜索工具... (空格聚焦，ESC 取消，回车执行第一个)"
            prefix-icon="Search"
            clearable
            @input="filterTools"
            @keyup.esc="clearSearch"
            @keyup.enter="executeFirstResult"
          />
          <el-button 
            class="add-button"
            type="primary" 
            :icon="Plus" 
            @click="showAddToolDialog"
          >
            添加
          </el-button>
        </div>
      </div>

      <!-- 工具列表 -->
      <div class="categories">
        <div
          v-for="category in filteredCategories"
          :key="category.Name"
          class="category-section"
        >
          <div class="category-header">
            <h2 class="category-title">{{ category.Name }}</h2>
            <div class="category-actions">
              <span class="tool-count">{{ category.Tool.length }} 个工具</span>
              <el-button 
                type="danger" 
                size="small" 
                icon="Delete"
                circle
                @click="confirmDeleteCategory(category.Name)"
              />
            </div>
          </div>
          <div class="tools-grid">
            <draggable
              v-model="category.Tool"
              :animation="150"
              ghost-class="ghost"
              @end="onDragEnd(category)"
              item-key="Name"
              :component-data="{ tag: 'div', class: 'tools-grid-inner' }"
            >
              <template #item="{ element: tool }">
                <el-card
                  class="tool-card"
                  shadow="hover"
                  @click="executeTool(tool)"
                  @contextmenu="showToolMenu($event, tool, category.Name)"
                  v-bind="getTooltipProps(tool)"
                >
                  <div class="tool-name">{{ tool.Name }}</div>
                </el-card>
              </template>
            </draggable>
          </div>
        </div>
      </div>
      
      <!-- 自定义工具提示 -->
      <div 
        v-if="tooltip.visible" 
        class="custom-tooltip"
        :class="{ 'tooltip-bottom': tooltip.direction === 'bottom' }"
        :style="{
          left: `${tooltip.x}px`,
          top: `${tooltip.y}px`,
          transform: tooltip.direction === 'top' ? 'translate(-50%, -100%)' : 'translate(-50%, 0)'
        }"
      >
        {{ tooltip.content }}
      </div>

      <!-- 添加工具对话框 -->
      <el-dialog
        v-model="addToolDialogVisible"
        title="添加工具"
        width="500px"
        class="add-tool-dialog"
        center
        append-to-body
      >
        <el-form :model="newTool" label-width="80px">
          <el-form-item label="工具名称">
            <el-input v-model="newTool.Name" placeholder="请输入工具名称" />
          </el-form-item>
          
          <el-form-item label="工具路径">
            <el-input v-model="newTool.Path" placeholder="请输入工具路径">
              <template #append>
                <el-button @click="selectToolPath">选择路径</el-button>
              </template>
            </el-input>
          </el-form-item>
          
          <el-form-item label="文件名">
            <el-input v-model="newTool.FileName" placeholder="请输入文件名" />
          </el-form-item>
          
          <el-form-item label="执行类型">
            <el-select v-model="newTool.Value" placeholder="请选择执行类型">
              <el-option
                v-for="type in ['Java8', 'Java11', 'Java17','Open', 'openterm']"
                :key="type"
                :label="type"
                :value="type"
              />
            </el-select>
          </el-form-item>
          
          <el-form-item label="命令">
            <el-input 
              v-model="newTool.Command" 
              placeholder="可选：自定义命令，默认-jar" 
              :default-value="-jar" 
            />
          </el-form-item>
          
          <el-form-item label="可选参数">
            <el-input v-model="newTool.Optional" placeholder="可选：命令行参数" />
          </el-form-item>
          
          <el-form-item label="工具描述">
            <el-input 
              v-model="newTool.Description" 
              type="textarea" 
              :rows="3"
              placeholder="请输入工具描述（可选）" 
            />
          </el-form-item>
          
          <el-form-item label="所属分类">
            <el-select 
              v-model="selectedCategory"
              placeholder="请选择分类"
              allow-create
              filterable
              default-first-option
            >
              <el-option
                v-for="category in categories.Category"
                :key="category.Name"
                :label="category.Name"
                :value="category.Name"
              />
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="addToolDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="addTool">确定</el-button>
          </span>
        </template>
      </el-dialog>

      <!-- 命令输出区域 -->
      <el-card class="output-card" v-if="outputText">
        <template #header>
          <div class="output-header">
            <span>命令输出</span>
            <el-button type="text" @click="outputText = ''">清除</el-button>
          </div>
        </template>
        <pre class="output-content">{{ outputText }}</pre>
      </el-card>

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

      <!-- 修改工具对话框 -->
      <el-dialog
        title="修改工具"
        v-model="editDialog.visible"
        width="500px"
        class="add-tool-dialog"
        center
        append-to-body
      >
        <el-form :model="editDialog.tool" label-width="80px">
          <el-form-item label="工具名称">
            <el-input v-model="editDialog.tool.Name" placeholder="请输入工具名称" />
          </el-form-item>
          
          <el-form-item label="工具路径">
            <el-input v-model="editDialog.tool.Path" placeholder="请输入工具路径">
              <template #append>
                <el-button @click="selectEditToolPath">选择路径</el-button>
              </template>
            </el-input>
          </el-form-item>
          
          <el-form-item label="文件名">
            <el-input v-model="editDialog.tool.FileName" placeholder="请输入文件名" />
          </el-form-item>
          
          <el-form-item label="执行类型">
            <el-select v-model="editDialog.tool.Value" placeholder="请选择执行类型">
              <el-option
                v-for="type in ['Java8', 'Java11', 'Java17','Open', 'openterm']"
                :key="type"
                :label="type"
                :value="type"
              />
            </el-select>
          </el-form-item>
          
          <el-form-item label="命令">
            <el-input v-model="editDialog.tool.Command" placeholder="可选：自定义命令" />
          </el-form-item>
          
          <el-form-item label="可选参数">
            <el-input v-model="editDialog.tool.Optional" placeholder="可选：命令行参数" />
          </el-form-item>
          
          <el-form-item label="工具描述">
            <el-input 
              v-model="editDialog.tool.Description" 
              type="textarea" 
              :rows="3"
              placeholder="请输入工具描述（可选）" 
            />
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
                v-for="category in categories.Category"
                :key="category.Name"
                :label="category.Name"
                :value="category.Name"
              />
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
          <span class="dialog-footer">
            <el-button @click="editDialog.visible = false">取消</el-button>
            <el-button type="primary" @click="submitToolEdit">确定</el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import { ref, reactive, onMounted, onBeforeUnmount, nextTick } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Folder, Edit, Delete, Plus } from '@element-plus/icons-vue'
import draggable from 'vuedraggable'

export default {
  name: 'App',
  components: {
    draggable,
    Folder,
    Edit,
    Delete,
    Plus
  },
  setup() {
    const categories = ref({ Category: [] });
    const filteredCategories = ref([]);
    const searchQuery = ref('');
    const showAddDialog = ref(false);
    const activeCategories = ref([]);
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
      tool: {
        Name: '',
        Path: '',
        FileName: '',
        Value: '',
        Command: '',
        Optional: '',
        Description: '' // 添加描述字段
      },
      category: '',
      originalName: ''
    });
    const toolTypes = ref([]);
    const searchInput = ref(null);
    const addToolDialogVisible = ref(false);
    const silentUpdate = ref(false);
    
    // 提示相关变量
    const tooltip = reactive({
      visible: false,
      content: '',
      x: 0,
      y: 0,
      direction: 'top' // 添加方向属性，用于确定提示显示在元素上方还是下方
    });

    // 获取工具提示属性
    const getTooltipProps = (tool) => {
      if (!tool.Description) {
        return {}; // 没有描述时不显示提示
      }
      
      return {
        'data-tooltip': tool.Description,
        'onMouseenter': (event) => showTooltip(event, tool.Description),
        'onMouseleave': hideTooltip
      };
    };

    // 显示提示 - 改进版本
    const showTooltip = (event, content) => {
      const rect = event.currentTarget.getBoundingClientRect();
      tooltip.content = content;
      
      // 获取滚动位置
      const scrollX = window.scrollX || document.documentElement.scrollLeft;
      const scrollY = window.scrollY || document.documentElement.scrollTop;
      
      // 基于元素在视口中的位置计算位置
      tooltip.x = rect.left + (rect.width / 2) + scrollX;
      
      // 计算可用空间
      const tooltipHeight = 40; // 估计的提示框高度（包括内边距和箭头）
      const tooltipWidth = 280; // 最大提示框宽度（来自CSS）
      const viewportHeight = window.innerHeight;
      const viewportWidth = window.innerWidth;
      
      // 检查元素上方是否有足够空间
      const spaceAbove = rect.top;
      const spaceBelow = viewportHeight - rect.bottom;
      
      // 根据可用空间确定方向
      if (spaceAbove < tooltipHeight && spaceBelow >= tooltipHeight) {
        // 上方空间不足但下方空间足够 - 在下方显示
        tooltip.direction = 'bottom';
        tooltip.y = rect.bottom + scrollY;
      } else {
        // 默认在上方显示
        tooltip.direction = 'top';
        tooltip.y = rect.top + scrollY;
      }
      
      // 确保提示框不会水平超出屏幕
      if (rect.left + (rect.width / 2) - (tooltipWidth / 2) < 0) {
        // 太靠近左边缘
        tooltip.x = tooltipWidth / 2 + scrollX + 10;
      } else if (rect.left + (rect.width / 2) + (tooltipWidth / 2) > viewportWidth) {
        // 太靠近右边缘
        tooltip.x = viewportWidth - (tooltipWidth / 2) + scrollX - 10;
      }
      
      tooltip.visible = true;
    };

    // 隐藏提示
    const hideTooltip = () => {
      tooltip.visible = false;
    };

    // 加载分类和工具列表
    const loadCategories = async () => {
      try {
        const result = await window.go.main.App.GetCategories();
        categories.value = result;
        filteredCategories.value = result.Category;
      } catch (err) {
        ElMessage.error(`加载工具列表失败: ${err}`);
      }
    };

    // 过滤工具
    const filterTools = () => {
      if (!searchQuery.value) {
        // 如果没有搜索查询，直接使用原始分类数据
        filteredCategories.value = categories.value.Category;
        return;
      }

      const query = searchQuery.value.toLowerCase();
      filteredCategories.value = categories.value.Category
        .map(category => {
          const filteredTools = category.Tool
            .filter(tool => {
              const nameMatch = tool.Name.toLowerCase().includes(query);
              const pathMatch = tool.Path.toLowerCase().includes(query);
              const descMatch = tool.Description ? tool.Description.toLowerCase().includes(query) : false;
              return nameMatch || pathMatch || descMatch;
            })
            .sort((a, b) => {
              // 首先按照打开次数排序
              if (b.OpenCount !== a.OpenCount) {
                return b.OpenCount - a.OpenCount;
              }
              // 然后按照名称匹配度排序
              const aExactMatch = a.Name.toLowerCase() === query;
              const bExactMatch = b.Name.toLowerCase() === query;
              if (aExactMatch && !bExactMatch) return -1;
              if (!aExactMatch && bExactMatch) return 1;
              const aStartMatch = a.Name.toLowerCase().startsWith(query);
              const bStartMatch = b.Name.toLowerCase().startsWith(query);
              if (aStartMatch && !bStartMatch) return -1;
              if (!aStartMatch && bStartMatch) return 1;
              return 0;
            });

          return {
            ...category,
            Tool: filteredTools
          };
        })
        .filter(category => category.Tool.length > 0);
    };

    // 执行工具
    const executeTool = async (tool) => {
      try {
        await window.go.main.App.ExecuteCommand(
          tool.Path,
          tool.Optional,
          tool.Value,
          tool.FileName
        );
      } catch (err) {
        ElMessage.error(`执行命令失败: ${err}`);
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

    // 显示添加工具对话框
    const showAddToolDialog = () => {
      addToolDialogVisible.value = true;
      resetNewToolForm();
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
          `确定要删除工具 "${contextMenu.selectedTool.Name}" 吗？`,
          '删除确认',
          {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }
        );

        await window.go.main.App.DeleteTool(
          contextMenu.selectedTool.Name,
          contextMenu.selectedCategory
        );
        
        ElMessage.success('删除成功');
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
        originalName: contextMenu.selectedTool.Name
      });
    };

    // 提交编辑后的工具
    const submitToolEdit = async () => {
      try {
        await window.go.main.App.UpdateTool(
          editDialog.originalName,
          editDialog.category,
          editDialog.tool
        );
      } catch (err) {
        ElMessage.error(`修改工具失败: ${err}`);
      }
    };

    // 添加工具
    const addTool = async () => {
      try {
        await window.go.main.App.AddTool(newTool, selectedCategory.value);
        addToolDialogVisible.value = false;
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
            type: 'warning'
          }
        );
        
        await deleteCategory(categoryName);
      } catch (err) {
        if (err !== 'cancel') {
          ElMessage.error(`删除失败: ${err}`);
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
      } else if (event.code === 'Enter') {
        executeFirstResult();
      }
    };

    // 执行第一个搜索结果
    const executeFirstResult = () => {
      // 如果有搜索结果
      if (filteredCategories.value.length > 0) {
        const firstCategory = filteredCategories.value[0];
        if (firstCategory.Tool && firstCategory.Tool.length > 0) {
          const firstTool = firstCategory.Tool[0];
          executeTool(firstTool);
          // 可选：清除搜索
          clearSearch();
          return;
        }
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

    // 拖拽结束
    const onDragEnd = async (category) => {
      try {
        silentUpdate.value = true;
        await window.go.main.App.UpdateCategoryTools(category.Name, category.Tool);
        ElMessage.success('工具顺序已更新');
        await loadCategories();
      } catch (err) {
        ElMessage.error(`更新工具顺序失败: ${err}`);
      } finally {
        silentUpdate.value = false;
      }
    };

    onMounted(async () => {
      await loadCategories();
      
      // 监听命令输出
      window.runtime.EventsOn('command-output', (output) => {
        outputText.value = output;
      });

      // 监听工具添加成功事件
      window.runtime.EventsOn('tool-added', () => {
        loadCategories();
        showAddDialog.value = false;
        ElMessage.success('工具添加成功');
      });

      // 监听工具更新成功事件
      window.runtime.EventsOn('tool-updated', () => {
        if (silentUpdate.value) return;
        loadCategories();
        editDialog.visible = false;
        ElMessage.success('工具修改成功');
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
      newTool,
      selectedCategory,
      outputText,
      contextMenu,
      editDialog,
      toolTypes,
      searchInput,
      addToolDialogVisible,
      silentUpdate,
      tooltip,
      loadCategories,
      filterTools,
      executeTool,
      loadToolTypes,
      resetNewToolForm,
      showAddToolDialog,
      selectToolPath,
      showToolMenu,
      closeContextMenu,
      openToolDirectory,
      deleteSelectedTool,
      showEditDialog,
      submitToolEdit,
      addTool,
      selectEditToolPath,
      confirmDeleteCategory,
      deleteCategory,
      handleGlobalKeydown,
      executeFirstResult,
      clearSearch,
      handleSpaceSearch,
      onDragEnd,
      getTooltipProps,
      showTooltip,
      hideTooltip,
      Plus,
      Folder,
      Edit,
      Delete
    };
  }
}
</script>

<style>
:root {
  --system-font: -apple-system, BlinkMacSystemFont, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "WenQuanYi Micro Hei", sans-serif;
}

body, 
.app-wrapper,
.container,
.category-title,
.tool-count,
.menu-item,
.el-button,
.el-input {
  font-family: var(--system-font);
}

.app-wrapper {
  position: relative;
  min-height: 100vh;
  transform-style: flat;
  background-color: transparent;
}

.container {
  position: relative;
  z-index: 1;
  transform: translateZ(0);
  background-color: transparent;
  width: 100%;
  max-width: 100%;
  margin: 0 auto;
  padding: 0 20px;
  box-sizing: border-box;
}

/* 顶部搜索栏 */
.header {
  margin-bottom: 24px;
}

.search-bar {
  display: flex;
  gap: 12px;
  align-items: center;
}

.search-bar .el-input {
  flex: 1;
}

.search-bar .el-input :deep(.el-input__wrapper) {
  background-color: rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 6px;
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
}

.search-bar .el-input :deep(.el-input__inner) {
  color: #ffffff;
  height: 30px;
  font-size: 13px;
}

.search-bar .el-input :deep(.el-input__prefix-icon) {
  color: rgba(255, 255, 255, 0.7);
  font-size: 14px;
}

.add-button {
  height: 30px;
  padding: 0 12px;
  font-size: 12px;
  border-radius: 6px;
  background: rgba(255, 255, 255, 0.2);
  color: #ffffff;
  font-weight: 500;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.add-button:hover {
  background: rgba(255, 255, 255, 0.3);
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
}

.add-button:active {
  transform: translateY(0);
  background: rgba(255, 255, 255, 0.2);
  box-shadow: none;
}

/* 搜索框交互优化 */
.search-bar .el-input :deep(.el-input__wrapper:hover),
.search-bar .el-input :deep(.el-input__wrapper.is-focus) {
  background-color: rgba(0, 0, 0, 0.15);
  border-color: rgba(255, 255, 255, 0.4);
}

/* 工具分类列表 */
.categories {
  margin-top: 20px;
}

.category-section {
  margin-bottom: 32px;
}

.category-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  padding-bottom: 8px;
  border-bottom: 2px solid rgba(255, 255, 255, 0.2);
}

.category-title {
  font-size: 18px;
  font-weight: 600;
  color: #ffffff;
  text-shadow: 0 1px 3px rgba(0, 0, 0, 0.3);
}

.tool-count {
  font-size: 13px;
  background-color: rgba(255, 255, 255, 0.2);
  color: #ffffff;
  padding: 4px 10px;
  border-radius: 10px;
}

/* 工具网格布局 - 修改为默认每行5个等宽按钮 */
.tools-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  padding: 12px 0;
  width: 100%;
}

.tools-grid-inner {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  width: 100%;
}

.tool-card {
  width: calc(20% - 9.6px); /* 一行5个等宽按钮，考虑间隔 */
  height: 48px; /* 统一高度 */
  display: flex;
  justify-content: center;
  align-items: center;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 6px;
  cursor: pointer;
  user-select: none;
  font-size: 14px;
  font-weight: 500;
  color: #fff;
  transition: all 0.2s;
  text-align: center;
  box-sizing: border-box;
  padding: 0 12px;
  margin: 0; /* 移除外边距 */
  flex-grow: 0; /* 不允许增长 */
  flex-shrink: 0; /* 不允许缩小 */
  font-family: var(--system-font); /* 使用系统字体 */
}

/* 自定义提示样式 */
.custom-tooltip {
  position: fixed;
  background: rgba(0, 0, 0, 0.8);
  color: #fff;
  padding: 8px 12px;
  border-radius: 4px;
  font-size: 13px;
  z-index: 9999;
  max-width: 280px;
  word-break: break-word;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.2);
  pointer-events: none;
  transition: opacity 0.2s;
  margin-top: -8px;
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
  border-top: 6px solid rgba(0, 0, 0, 0.8);
}

/* 添加下方提示样式 */
.tooltip-bottom::after {
  bottom: auto;
  top: -6px;
  border-top: none;
  border-bottom: 6px solid rgba(0, 0, 0, 0.8);
}

/* 响应式调整 - 调整默认显示为4个按钮 */
@media (min-width: 1401px) {
  .tool-card {
    width: calc(16.666% - 10px); /* 大屏幕一行6个等宽按钮 */
  }
}

@media (max-width: 1100px) {
  .tool-card {
    width: calc(25% - 9px); /* 一行4个等宽按钮 */
  }
}

@media (max-width: 768px) {
  .tool-card {
    width: calc(33.333% - 8px); /* 一行3个等宽按钮 */
  }
}

@media (max-width: 480px) {
  .tool-card {
    width: calc(50% - 6px); /* 一行2个等宽按钮 */
  }
}

.tool-name {
  font-size: 13px;
  font-weight: 500;
  color: rgba(255, 255, 255, 0.9);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  width: 100%;
  max-width: 100%;
  text-align: center;
  line-height: 1.3;
  margin: 0 auto;
  padding: 0;
  font-family: var(--system-font); /* 使用系统字体 */
}

/* 工具卡片交互效果 */
.tool-card:hover {
  background: rgba(255, 255, 255, 0.12);
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.tool-card:active {
  transform: scale(0.98);
  background: rgba(255, 255, 255, 0.08);
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

/* 添加分隔线样式 */
.menu-divider {
  height: 1px;
  background: rgba(255, 255, 255, 0.1);
  margin: 4px 0;
}


::-webkit-scrollbar {
    width: 0 !important;
}
::-webkit-scrollbar {
    width: 0 !important;height: 0;
}

.content-wrapper {
  -ms-overflow-style: none;
  scrollbar-width: none;
  overflow-y: auto;
}

.content-wrapper::-webkit-scrollbar {
  display: none;
}

.category-actions {
  display: flex;
  align-items: center;
  gap: 12px;
}

.category-actions .el-button {
  margin-left: 8px;
  opacity: 0.7;
  transition: all 0.2s ease;
}

.category-actions .el-button:hover {
  opacity: 1;
  transform: scale(1.1);
}

.category-header:hover .category-actions .el-button {
  opacity: 1;
}

/* 添加容器的焦点样式 */
.container {
  outline: none; /* 移除默认的焦点轮廓 */
}

/* 优化搜索框的焦点样式 */
.search-bar .el-input :deep(.el-input__wrapper.is-focus) {
  box-shadow: 0 0 0 1px var(--el-color-primary) inset;
  border-color: var(--el-color-primary);
  background-color: rgba(0, 0, 0, 0.2);
}

.category {
    font-family: var(--system-font) !important;
}

/* 拖拽样式 */
.ghost {
  opacity: 0.5;
  background: rgba(0, 123, 255, 0.2) !important;
  border: 2px dashed #0d6efd !important;
}
</style>

