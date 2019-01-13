import React, { Component, Fragment } from 'react';
import { formatMessage, FormattedMessage } from 'umi/locale';
import { Form, Input, Upload, Button } from 'antd';
import { connect } from 'dva';
import PageHeaderWrapper from '@/components/PageHeaderWrapper';
import GridContent from '@/components/PageHeaderWrapper/GridContent';
import styles from './UploadVideo.less';
// import { getTimeDistance } from '@/utils/utils';

const FormItem = Form.Item;

const UploadView = () => (
  <Fragment>
    <div className={styles.avatar_title}>
      <FormattedMessage id="app.settings.basic.avatar" defaultMessage="Avatar" />
    </div>
    <Upload fileList={[]}>
      <div className={styles.button_view}>
        <Button icon="upload">
          <FormattedMessage id="app.settings.basic.change-avatar" defaultMessage="Change avatar" />
        </Button>
      </div>
    </Upload>
  </Fragment>
);

@Form.create()
class UploadVideo extends Component {
  componentDidMount() {
  }

  getViewDom = ref => {
    this.view = ref;
  };

  render() {
    const {
      form: { getFieldDecorator },
    } = this.props;
    return (
      <PageHeaderWrapper title="上传课程视频">
        <GridContent>
          <div className={styles.baseView} ref={this.getViewDom}>
            <div className={styles.left}>
              <Form onSubmit={this.handleSubmit}>
                <FormItem label={formatMessage({ id: 'course.upload.title' })}>
                  {getFieldDecorator('name', {
                    rules: [
                      {
                        required: true,
                        message: formatMessage({ id: 'course.upload.title-message' }, {}),
                      },
                    ],
                  })(<Input placeholder={formatMessage({ id: 'course.upload.title-placeholder' })} />)}
                </FormItem>
                <FormItem label={formatMessage({ id: 'course.upload.abstract' })}>
                  {getFieldDecorator('profile', {
                    rules: [
                      {
                        required: true,
                        message: formatMessage({ id: 'course.upload.abstract-message' }, {}),
                      },
                    ],
                  })(
                    <Input.TextArea
                      placeholder={formatMessage({ id: 'course.upload.abstract-placeholder' })}
                      rows={4}
                    />
                  )}
                </FormItem>
                <Button type="primary">
                  <FormattedMessage
                    id="app.settings.basic.update"
                    defaultMessage="Update Information"
                  />
                </Button>
              </Form>
            </div>
            <div className={styles.right}>
              <UploadView />
            </div>
          </div>
        </GridContent>
      </PageHeaderWrapper>
    );
  }
}

export default UploadVideo;
