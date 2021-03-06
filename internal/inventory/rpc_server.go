// Code generated, do not edit
package inventory

// Server
type Server struct {
	UnimplementedInventoryServer
	inventory Inventory
}

// NewServer is used to return a new Server, which implements the inventory RPC server.
func NewLDAPServer(invClient Inventory) (*Server, error) {
	return &Server{
		inventory: invClient,
	}, nil
}

// .File.Service Inventory
func (e *Server) ListNetworkHosts(req *Empty, stream Inventory_ListNetworkHostsServer) error {
	// .Method.Name ListNetworkHosts
	// .Method.InputType .inventory.Empty
	// .Method.OutputType .inventory.NetworkHost

	results, err := e.inventory.ListNetworkHosts()
	if err != nil {
		return err
	}

	// outputType: NetworkHost
	if results != nil {
		for _, r := range results {
			xxx := &NetworkHost{}
			xxx.Role = r.Role                       // TYPE_STRING
			xxx.Group = r.Group                     // TYPE_STRING
			xxx.Name = r.Name                       // TYPE_STRING
			xxx.OperatingSystem = r.OperatingSystem // TYPE_STRING
			xxx.Platform = r.Platform               // TYPE_STRING
			xxx.Type = r.Type                       // TYPE_STRING
			xxx.Domain = r.Domain                   // TYPE_STRING
			xxx.Description = r.Description         // TYPE_STRING
			// Unhandled Type
			// name: watch
			// type: TYPE_BOOL
			// label: LABEL_OPTIONAL
			if r.InetAddress != nil {
				xxx.InetAddress = r.InetAddress // TYPE_STRING
			}
			if r.Inet6Address != nil {
				xxx.Inet6Address = r.Inet6Address // TYPE_STRING
			}
			if r.MacAddress != nil {
				xxx.MacAddress = r.MacAddress // TYPE_STRING
			}
			// Unhandled Type
			// name: last_seen
			// type: TYPE_MESSAGE
			// label: LABEL_OPTIONAL
			xxx.Dn = r.Dn // TYPE_STRING

			if err := stream.Send(xxx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *Server) ListNetworkIDs(req *Empty, stream Inventory_ListNetworkIDsServer) error {
	// .Method.Name ListNetworkIDs
	// .Method.InputType .inventory.Empty
	// .Method.OutputType .inventory.NetworkID

	results, err := e.inventory.ListNetworkIDs()
	if err != nil {
		return err
	}

	// outputType: NetworkID
	if results != nil {
		for _, r := range results {
			xxx := &NetworkID{}
			xxx.Name = r.Name // TYPE_STRING
			if r.MacAddress != nil {
				xxx.MacAddress = r.MacAddress // TYPE_STRING
			}
			if r.IpAddress != nil {
				xxx.IpAddress = r.IpAddress // TYPE_STRING
			}
			if r.ReportingSource != nil {
				xxx.ReportingSource = r.ReportingSource // TYPE_STRING
			}
			if r.ReportingSourceInterface != nil {
				xxx.ReportingSourceInterface = r.ReportingSourceInterface // TYPE_STRING
			}
			// Unhandled Type
			// name: last_seen
			// type: TYPE_MESSAGE
			// label: LABEL_OPTIONAL
			xxx.Dn = r.Dn // TYPE_STRING

			if err := stream.Send(xxx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (e *Server) ListZigbeeDevices(req *Empty, stream Inventory_ListZigbeeDevicesServer) error {
	// .Method.Name ListZigbeeDevices
	// .Method.InputType .inventory.Empty
	// .Method.OutputType .inventory.ZigbeeDevice

	results, err := e.inventory.ListZigbeeDevices()
	if err != nil {
		return err
	}

	// outputType: ZigbeeDevice
	if results != nil {
		for _, r := range results {
			xxx := &ZigbeeDevice{}
			xxx.Name = r.Name               // TYPE_STRING
			xxx.Description = r.Description // TYPE_STRING
			xxx.Dn = r.Dn                   // TYPE_STRING
			// Unhandled Type
			// name: last_seen
			// type: TYPE_MESSAGE
			// label: LABEL_OPTIONAL
			xxx.IotZone = r.IotZone                   // TYPE_STRING
			xxx.Type = r.Type                         // TYPE_STRING
			xxx.SoftwareBuildId = r.SoftwareBuildId   // TYPE_STRING
			xxx.DateCode = r.DateCode                 // TYPE_STRING
			xxx.Model = r.Model                       // TYPE_STRING
			xxx.Vendor = r.Vendor                     // TYPE_STRING
			xxx.ManufacturerName = r.ManufacturerName // TYPE_STRING
			xxx.PowerSource = r.PowerSource           // TYPE_STRING
			xxx.ModelId = r.ModelId                   // TYPE_STRING

			if err := stream.Send(xxx); err != nil {
				return err
			}
		}
	}

	return nil
}

// func (r *Server) ListNetworkHosts_example(ctx context.Context, request *Empty) (*SearchResponse, error) {
// 	response := &SearchResponse{}
//
// 	hosts, err := r.inventory.ListNetworkHosts()
// 	if err != nil {
// 		return response, err
// 	}
//
// 	if hosts != nil {
// 		for _, h := range *hosts {
// 			host := &NetworkHost{
// 				Description:     h.Description,
// 				Dn:              h.Dn,
// 				Domain:          h.Domain,
// 				Group:           h.Group,
// 				Name:            h.Name,
// 				OperatingSystem: h.OperatingSystem,
// 				Platform:        h.Platform,
// 				Role:            h.Role,
// 				Type:            h.Type,
// 			}
//
// 			if h.InetAddress != nil {
// 				host.InetAddress = h.InetAddress
// 			}
//
// 			if h.Inet6Address != nil {
// 				host.Inet6Address = h.Inet6Address
// 			}
//
// 			if h.MacAddress != nil {
// 				host.MacAddress = h.MacAddress
// 			}
// 			if h.LastSeen != nil {
// 				lastSeen, err := ptypes.TimestampProto(h.LastSeen)
// 				if err != nil {
// 					log.Error(err)
// 				}
// 				host.LastSeen = lastSeen
// 			}
//
// 			response.Hosts = append(response.Hosts, host)
// 		}
// 	}
//
// 	return response, nil
// }
