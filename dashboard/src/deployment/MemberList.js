import React, { Component } from 'react';
import styled from 'react-emotion';
import { Accordion, Header, Icon, List, Segment } from 'semantic-ui-react';
import CommandInstruction from '../util/CommandInstruction.js';

const Field = styled('div')`
  padding-top: 0.3em;
  padding-bottom: 0.3em;
`;

const FieldLabel = styled('span')`
  width: 5rem;
  display: inline-block;
`;

const FieldIcons = styled('div')`
  float: right;
`;

const MemberListView = ({group, activeMemberID, onClick, members, namespace}) => (
  <Segment>
    <Header>{group}</Header>
    <List divided>
      {members.map((item) => <MemberView key={item.id} memberInfo={item} active={item.id === activeMemberID} onClick={onClick} namespace={namespace}/>)}
    </List>
  </Segment>
);

const MemberView = ({memberInfo, namespace, active, onClick}) => (
  <List.Item>
    <Accordion>
      <Accordion.Title active={active} onClick={() => onClick(memberInfo.id)}>
        <Icon name='dropdown' /> {memberInfo.id}
      </Accordion.Title>
      <Accordion.Content active={active}>
        <Field>
          <FieldLabel>Pod:</FieldLabel> 
          {memberInfo.pod_name || "-"}
          {(memberInfo.pod_name) ?
            <FieldIcons>
              <CommandInstruction 
                trigger={<Icon link name="file outline alternate"/>}
                command={createLogPodCommand(memberInfo.pod_name, namespace)}
                title="Get Pod Logs"
                description="To get the log output of this pod, run:"
              />
              <CommandInstruction 
                trigger={<Icon link name="zoom"/>}
                command={createDescribePodCommand(memberInfo.pod_name, namespace)}
                title="Describe Pod Information"
                description="To get more information on the state of this pod, run:"
              />
            </FieldIcons>
          : null}
        </Field>
        <Field>
          <FieldLabel>PVC:</FieldLabel> 
          {memberInfo.pvc_name || "-"}
          <FieldIcons>
            {(memberInfo.pvc_name) ?
              <CommandInstruction 
                trigger={<Icon link name="zoom"/>}
                command={createDescribePvcCommand(memberInfo.pvc_name, namespace)}
                title="Describe PersistentVolumeClaim Information"
                description="To get more information on the state of this PVC, run:"
              />
            : null}
          </FieldIcons>
        </Field>
        <Field>
          <FieldLabel>PV:</FieldLabel>
          {memberInfo.pv_name || "-"}
          <FieldIcons>
            {(memberInfo.pv_name) ?
              <CommandInstruction 
                trigger={<Icon link name="zoom"/>}
                command={createDescribePvCommand(memberInfo.pv_name)}
                title="Describe PersistentVolume Information"
                description="To get more information on the state of this PV, run:"
              />
            : null}
          </FieldIcons>
        </Field>
      </Accordion.Content>
    </Accordion>
  </List.Item>
);

function createDescribePodCommand(podName, namespace) {
  return `kubectl describe pod -n ${namespace} ${podName}`;
}

function createLogPodCommand(podName, namespace) {
  return `kubectl logs -n ${namespace} ${podName}`;
}

function createDescribePvcCommand(pvcName, namespace) {
  return `kubectl describe pvc -n ${namespace} ${pvcName}`;
}

function createDescribePvCommand(pvName) {
  return `kubectl describe pv ${pvName}`;
}

class MemberList extends Component {
  state = {};

  onClick = (id) => { 
    this.setState({activeMemberID:(this.state.activeMemberID === id) ? null : id}); 
  }

  render() {
    return (<MemberListView 
      group={this.props.group} 
      members={this.props.members}
      activeMemberID={this.state.activeMemberID}
      onClick={this.onClick}
      namespace={this.props.namespace}
    />);
  }
}

export default MemberList;
