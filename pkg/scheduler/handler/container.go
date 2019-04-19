package handler

import (
	"github.com/streadway/amqp"
)

type DeploymentHandler struct{}

func (DeploymentHandler) Name() string {
	return "ContainerDeploymentRequest"
}

func (DeploymentHandler) Handle(m amqp.Delivery) error {
	//message := &protocol.ChangeContainerStateRequest{}
	//if err := proto.Unmarshal(m.Body, message); err != nil {
	//	return err
	//}
	//container, err := containers.FindContainerByID(context.Background(), message.Id)
	//if err != nil || container == nil {
	//	return err
	//}
	//service, err := docker.ServiceFindByName(container.ID)
	//if err != nil {
	//	return err
	//}
	//if service == nil {
	//	replicas := uint64(1)
	//	resources := &swarm.Resources{
	//		MemoryBytes: int64(1 * 1024 * 1024), // 1 == mo
	//		NanoCPUs:    int64(100000000 * 2),
	//	}
	//	response, err := docker.ServiceCreate(swarm.ServiceSpec{
	//		Annotations: swarm.Annotations{
	//			Name: container.ID,
	//			Labels: map[string]string{
	//				"traefik.enable": "true",
	//				"traefik.domain": "expected.sh",
	//				//"traefik.frontend.rule": "Host:hello.expected.sh",
	//				"traefik.port": "80",
	//				//"traefik.docker.network": "private",
	//			},
	//		},
	//		TaskTemplate: swarm.TaskSpec{
	//			ContainerSpec: swarm.ContainerSpec{
	//				Image: container.Image,
	//				Env:   convertEnv(container.Environment),
	//			},
	//			Resources: &swarm.ResourceRequirements{
	//				Limits:       resources,
	//				Reservations: resources,
	//			},
	//		},
	//		Mode: swarm.ServiceMode{
	//			Replicated: &swarm.ReplicatedService{
	//				Replicas: &replicas,
	//			},
	//		},
	//		Networks: []swarm.NetworkAttachmentConfig{
	//			{
	//				Target: "traefik-net",
	//			},
	//		},
	//	})
	//	if err != nil {
	//		return err
	//	}
	//	err = aws.Route53AddRecord("Z2MP7C8I98E8MT", "CNAME", container.ID+".expected.sh", []string{"prod.expected.sh"})
	//	if err != nil {
	//		return err
	//	}
	//	logrus.Infoln(response.ID)
	//} else {
	//
	//}
	return nil
}

type StartHandler struct{}

func (StartHandler) Name() string {
	return "ContainerStartRequest"
}

func (StartHandler) Handle(m amqp.Delivery) error {
	//message := &protocol.ContainerStartRequest{}
	//if err := proto.Unmarshal(m.Body, message); err != nil {
	//	return err
	//}
	//container, err := containers.FindContainerByID(context.Background(), message.Id)
	//if err != nil || container == nil {
	//	return err
	//}
	return nil
}

type StopHandler struct{}

func (StopHandler) Name() string {
	return "ContainerStopRequest"
}

func (StopHandler) Handle(m amqp.Delivery) error {
	//message := &protocol.ContainerStopRequest{}
	//if err := proto.Unmarshal(m.Body, message); err != nil {
	//	return err
	//}
	//container, err := containers.FindContainerByID(context.Background(), message.Id)
	//if err != nil || container == nil {
	//	return err
	//}
	return nil
}
