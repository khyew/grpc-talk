using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Grpc.Core;
using Demo;

namespace client
{
    class Program
    {
        static void Main(string[] args)
        {
            String[] arguments = Environment.GetCommandLineArgs();

            if (arguments.Length != 4) {
                Console.WriteLine("Usage: {0} IP PORT NAME", arguments[0]);
                return;
            }

            var ip = arguments[1];
            var port = arguments[2];
            var name = arguments[3];

            Channel channel = new Channel(ip + ":" + port, ChannelCredentials.Insecure);
            var client = new DemoService.DemoServiceClient(channel);

            var reply = client.Hello(new HelloRequest
            {
                Name = name
            });

            Console.WriteLine(reply.Reply);
        }
    }
}
