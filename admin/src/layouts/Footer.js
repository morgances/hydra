import React, { Fragment } from 'react';
import { Layout, Icon } from 'antd';
import GlobalFooter from '@/components/GlobalFooter';

const { Footer } = Layout;
const FooterView = () => (
  <Footer style={{ padding: 0 }}>
    <GlobalFooter
      links={[
        {
          key: '官方网站',
          title: '官方网站',
          href: 'https://www.smartestee.com/',
          blankTarget: true,
        },
        {
          key: 'Github',
          title: 'Github',
          href: 'https://github.com/fengyfei',
          blankTarget: true,
        },
        {
          key: 'DDR',
          title: '杨鼎睿',
          href: 'https://github.com/yhyddr',
          blankTarget: true,
        },
      ]}
      copyright={
        <Fragment>
          Copyright <Icon type="copyright" /> 2019 极智人（北京）科技有限公司
        </Fragment>
      }
    />
  </Footer>
);
export default FooterView;
