import Breadcrumb from './src/breadcrumb.mjs';
import BreadcrumbItem from './src/breadcrumb-item.mjs';
export { breadcrumbProps } from './src/breadcrumb2.mjs';
export { breadcrumbItemProps } from './src/breadcrumb-item2.mjs';
export { breadcrumbKey } from './src/constants.mjs';
import { withInstall, withNoopInstall } from '../../utils/vue/install.mjs';

const ElBreadcrumb = withInstall(Breadcrumb, {
  BreadcrumbItem
});
const ElBreadcrumbItem = withNoopInstall(BreadcrumbItem);

export { ElBreadcrumb, ElBreadcrumbItem, ElBreadcrumb as default };
//# sourceMappingURL=index.mjs.map
